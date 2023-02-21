import Head from "next/head";
import { useEffect, useRef, useState } from "react";

export default function Home() {
  // const [isTermInit, SetIsTettermInit] = useState(false);
  // const [termCommand, setTermCommand] = useState("");
  const termCommand = useRef("");
  let isTermInit = false;
  useEffect(() => {
    const initTerminal = async () => {
      const { Terminal } = await import("xterm");
      const { AttachAddon } = await import("xterm-addon-attach");
      const term = new Terminal();
      // Connect to the WebSocket endpoint.
      const socket = new WebSocket("ws://localhost:8080/terminal");
      const attachAddon = new AttachAddon(socket);

      // Attach the socket to term
      term.loadAddon(attachAddon);
      // Create the terminal.
      term.open(document.getElementById("terminal"));

      // // When the WebSocket connection is opened, send the initial message.
      // socket.addEventListener("open", function (event) {
      //   socket.send("Hello, world!");
      // });

      // // When a message is received from the server, write it to the terminal.
      // socket.addEventListener("message", function (event) {
      //   term.write(event.data + "\n\r");
      // });

      // // When the user types in the terminal, send the input to the server.
      // term.onKey(function (arg1, arg2) {
      //   let data = "hello";
      //   console.log(arg1);
      //   if (arg1.key == "\r") {
      //     console.log("TERM COMMAND", termCommand.current);
      //     socket.send(termCommand.current + "\r");
      //     return;
      //   }
      //   // setTermCommand((prev) => {
      //   //   console.log(data);
      //   //   return prev.concat(data);
      //   // });
      //   termCommand.current += arg1.key;
      // });
    };
    if (!isTermInit) {
      isTermInit = true;
      initTerminal();
    }
  }, []);
  // useEffect(() => {
  //   console.log("here: ", termCommand);
  // }, [setTermCommand]);
  return (
    <>
      <Head>
        <title>Terminal</title>
        <meta name="description" content="Examplte Terminal" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <div id="terminal"></div>
      </main>
    </>
  );
}
