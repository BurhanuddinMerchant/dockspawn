package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func main2() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3000/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "5000",
				},
			},
		},
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"sleep", "15"},
		Tty:   false,
		ExposedPorts: nat.PortSet{
			"3000/tcp": struct{}{},
		},
	}, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
func main2__() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	// defer reader.Close()
	// io.Copy(os.Stdout, reader)

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3000/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "5000",
				},
			},
		},
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "hello:latest",
		// Cmd:   []string{"sleep", "15"},
		Tty: false,
		ExposedPorts: nat.PortSet{
			"3000/tcp": struct{}{},
		},
	}, hostConfig, nil, nil, "new_container")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

// func buildFromDockerFile(image_tag string) {
// 	ctx := context.Background()
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		log.Fatal(err, " :unable to init client")
// 	}

// 	buf := new(bytes.Buffer)
// 	tw := tar.NewWriter(buf)
// 	defer tw.Close()

// 	dockerFile := "Dockerfile"
// 	dockerFileReader, err := os.Open("/home/burhanuddin/Dev/web/dock-spawn/go/Dockerfile")
// 	if err != nil {
// 		log.Fatal(err, " :unable to open Dockerfile")
// 	}
// 	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
// 	if err != nil {
// 		log.Fatal(err, " :unable to read dockerfile")
// 	}

// 	tarHeader := &tar.Header{
// 		Name: dockerFile,
// 		Size: int64(len(readDockerFile)),
// 	}
// 	err = tw.WriteHeader(tarHeader)
// 	if err != nil {
// 		log.Fatal(err, " :unable to write tar header")
// 	}
// 	_, err = tw.Write(readDockerFile)
// 	if err != nil {
// 		log.Fatal(err, " :unable to write tar body")
// 	}
// 	dockerFileTarReader := bytes.NewReader(buf.Bytes())

// 	imageBuildResponse, err := cli.ImageBuild(
// 		ctx,
// 		dockerFileTarReader,
// 		types.ImageBuildOptions{
// 			Context:    dockerFileTarReader,
// 			Dockerfile: dockerFile,
// 			Tags:       []string{image_tag},
// 			Remove:     true})
// 	if err != nil {
// 		log.Fatal(err, " :unable to build docker image")
// 	}
// 	defer imageBuildResponse.Body.Close()
// 	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
// 	if err != nil {
// 		log.Fatal(err, " :unable to read image build response")
// 	}
// 	main2__()
// }
func buildFromDockerFile() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFile := "Dockerfile"
	dockerFileReader, err := os.Open("/home/burhanuddin/Dev/web/dock-spawn/go/Dockerfile")
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, " :unable to read dockerfile")
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Dockerfile: dockerFile,
			Tags:       []string{"hello"},
			Remove:     true})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
	main2__()
}
func buildFromStringDockerFile() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	dockerfile := `
	FROM alpine
	RUN echo "Hello from a Dockerfile!"
	`

	buildContext := strings.NewReader(dockerfile)

	buildOptions := types.ImageBuildOptions{
		Tags:   []string{"myimage:latest"},
		Remove: true,
	}

	response, err := cli.ImageBuild(ctx, buildContext, buildOptions)
	if err != nil {
		log.Fatalf("Error building Docker image: %v", err)
	}

	defer response.Body.Close()

	containerConfig := &container.Config{
		Image: "myimage:latest",
		Cmd:   []string{"sh", "-c", "echo Hello from a container!"},
	}
	hostConfig := &container.HostConfig{}
	networkingConfig := &network.NetworkingConfig{}

	container, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkingConfig, nil, "")
	if err != nil {
		log.Fatalf("Error creating Docker container: %v", err)
	}

	if err := cli.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("Error starting Docker container: %v", err)
	}

	log.Printf("Container ID: %s", container.ID)
}
func main_() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
	buildFromStringDockerFile()
}

func terminal_process(container_id string, ws websocket.Conn) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	term_size := [2]uint{80, 24}
	exec_cfg := types.ExecConfig{User: "", Privileged: false, Tty: false, ConsoleSize: &term_size, AttachStdin: true, AttachStdout: true, AttachStderr: true, Detach: true, DetachKeys: "", Env: nil, WorkingDir: "/", Cmd: []string{"sh", "-c"}}
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}
	exec_config := types.ExecStartCheck{Detach: false, Tty: true, ConsoleSize: &term_size}
	exec_id, err := cli.ContainerExecCreate(ctx, container_id, exec_cfg)
	if err != nil {
		fmt.Print("Container Exec Create failed")
		return
	}
	log.Println("Container Exec Create success")
	// err = cli.ContainerExecStart(ctx, exec_id.ID, exec_config)
	log.Println("hkjfhjsajfkhdsfhkasjdfhjksdhf")
	hresp, err := cli.ContainerExecAttach(ctx, exec_id.ID, exec_config)

	log.Println(hresp.Conn.LocalAddr())
	log.Println(hresp.Conn.RemoteAddr())
	if err != nil {
		panic(err)
	}
	// defer hresp.Close()
	// use the response object to stream the output of the command
	log.Println(hresp.Conn)
	go func() {
		for {
			log.Println("buffer me: ", hresp.Reader.Buffered())
			// if hresp.Reader.Buffered() > 0 {

			var buf = make([]byte, 1024)
			n, err := hresp.Conn.Read(buf)

			fmt.Println("n: ", n)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Print("askljfjlskdjfklas", string(buf[:n]))
			if n != 0 {
				err = ws.WriteMessage(websocket.TextMessage, buf[:n])
				// }
				if err != nil {
					fmt.Println(err)
				}
			}
			// }
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("err***", err)
			log.Println("*****", msg)
			if len(msg) > 0 {
				fmt.Println(msg)
				bytes_written, err := hresp.Conn.Write(msg)

				if err != nil {
					log.Println(err)
				}
				log.Println(bytes_written)
			}
		}
	}()

}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/terminal", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("lskfjlskdfjlsjflksjdklfs")

		// TODO: Write logic for checking if container id has been sent from the frontend if not then start a new container
		// // Start a goroutine to handle incoming WebSocket messages.
		terminal_process("bb2b5f24febd", *conn)
		// _, msg, err := conn.ReadMessage()
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }

		// // Print the message to the console for debugging purposes.
		// log.Printf("Received message: %s", msg)

		// // Write the message back to the client.
		// if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		// 	log.Println(err)
		// 	return
		// }

	})
	fmt.Print("server up")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
