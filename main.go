package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"regexp"
)

func handleHookedRequest(w http.ResponseWriter, r *http.Request) {
	// Modify the request
	modifiedRequest, err := editRequest(r)
	if err != nil {
		log.Println("Failed modifying the request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("new request:", *modifiedRequest)

	// Send new request
	c := http.Client{}
	resp, err := c.Do(modifiedRequest)
	if err != nil {
		log.Println("The modified request failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("response:", resp)

	// Modify response
	modifiedResponse, err := editResponse(resp, modifiedRequest)
	if err != nil {
		log.Println("Failed modifying response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	modifiedResponse.Write(w)
}

func editRequest(r *http.Request) (*http.Request, error) {
	var empty *http.Request
	removeProxyHeaders(r)
	// Read the whole request
	reqBytes, err := httputil.DumpRequest(r, true)
	if err != nil {
		return empty, err
	}

	// Edit response content
	tempfile, err := edit(reqBytes, "fart-req")
	if err != nil {
		return empty, err
	}
	defer os.Remove(tempfile)

	// Read modified content
	modified, err := os.Open(tempfile)
	if err != nil {
		return empty, err
	}
	defer modified.Close()
	fileReader := bufio.NewReader(modified)
	newReq, err := http.ReadRequest(fileReader)
	if err != nil {
		return empty, err
	}

	// Reset file to the beginning to read first line (we need to parse the URL)
	modified.Seek(0, 0)
	line, err := fileReader.ReadString('\n')
	if err != nil {
		return empty, err
	}
	newReq.RequestURI = ""
	newURL, err := parseURL(line)
	newURL.Host = newReq.Host
	newURL.Scheme = r.URL.Scheme
	newReq.Proto = r.Proto
	log.Println("New url", *newURL)
	newReq.URL = newURL
	return newReq, err
}

func editResponse(resp *http.Response, req *http.Request) (*http.Response, error) {
	var empty *http.Response
	respBytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return empty, err
	}

	// Edit response content
	tempfile, err := edit(respBytes, "fart-resp")
	if err != nil {
		return empty, err
	}
	defer os.Remove(tempfile)

	// Read modified content
	modified, err := os.Open(tempfile)
	if err != nil {
		return empty, err
	}
	defer modified.Close()
	fileReader := bufio.NewReader(modified)

	return http.ReadResponse(fileReader, req)
}

func edit(content []byte, prefix string) (string, error) {
	// Create a tempfile and write the request to it
	tempfile, err := ioutil.TempFile("", prefix)
	if err != nil {
		return "", err
	}

	_, err = tempfile.Write(content)
	if err != nil {
		return "", err
	}
	tempfile.Close()

	// Start vim to modify the request
	if err := openEditor(tempfile.Name()); err != nil {
		return "", err
	}

	return tempfile.Name(), nil
}

func parseURL(line string) (*url.URL, error) {
	re := regexp.MustCompile("(GET|POST|PUT|PATCH|HEAD|DELETE|CONNECT) (.*) HTTP/(1.0|1.1|2.0)")
	vals := re.FindStringSubmatch(line)
	if len(vals) != 4 {
		return &url.URL{}, errors.New(fmt.Sprintf("cannot parse url: %s", line))
	}
	return url.ParseRequestURI(vals[2])
}

func removeProxyHeaders(r *http.Request) {
	r.RequestURI = ""
	r.Header.Del("Accept-Encoding")
	r.Header.Del("Proxy-Connection")
	r.Header.Del("Proxy-Authenticate")
	r.Header.Del("Proxy-Authorization")
	if r.Header.Get("Connection") == "close" {
		r.Close = false
	}
	r.Header.Del("Connection")
}

func openEditor(path string) error {
	editor := exec.Command("vim", path)
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr
	return editor.Run()
}

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handleHookedRequest),
	}
	log.Fatal(server.ListenAndServe())
}
