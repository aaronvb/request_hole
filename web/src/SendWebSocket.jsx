import { useState, useEffect } from "react";
import { useQuery, gql } from "@apollo/client";

export const SERVER_INFO = gql`
  query GetServerInfo {
    serverInfo {
      request_address
      request_port
      protocol
    }
  }
`;

function SendWebSocket(props) {
  const { data } = useQuery(SERVER_INFO);
  const [url, setUrl] = useState("");
  const [body, setBody] = useState(JSON.stringify({ hello: "world" }));
  const [connected, setConnected] = useState(false);
  const [connection, setConnection] = useState(null);

  const sendRequest = () => {
    connection.send(body);
  };

  const connect = () => {
    const socket = new WebSocket(url);
    socket.addEventListener("open", function (event) {
      setConnected(true);
      setConnection(socket);
    });

    socket.addEventListener("close", function (event) {
      setConnected(false);
      setConnection(null);
    });
  };

  const disconnect = () => {
    if (connection) {
      connection.close();
      setConnected(false);
    }
  };

  useEffect(() => {
    if (data) {
      setUrl(
        `${data.serverInfo.protocol}://${data.serverInfo.request_address}:${data.serverInfo.request_port}`
      );
    }
  }, [data]);

  if (!props.visible) {
    return <div></div>;
  } else {
    return (
      <section className="text-gray-600 bg-gray-100 body-font h-full">
        <div className="container p-5 mx-auto max-w-2xl">
          <div className="bg-white rounded shadow py-4 px-4">
            <h2 className="text-gray-900 text-lg mb-1 font-medium title-font">
              Send a WebSocket Message
            </h2>
            <div className="flex flex-wrap mb-4">
              <div className="w-full">
                <div className="relative">
                  <label
                    htmlFor="url"
                    className="tracking-midwest text-xs text-gray-400"
                  >
                    URL
                  </label>
                  {connected === false ? (
                    <input
                      type="text"
                      id="url"
                      name="url"
                      className="w-full rounded border border-gray-300 focus:border-red-500 focus:bg-white focus:ring-2 focus:ring-red-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
                      value={url}
                      onChange={(e) => setUrl(e.target.value)}
                    />
                  ) : (
                    <div className="text-green-500">Connected to {url}</div>
                  )}
                </div>
              </div>
            </div>
            {connected && (
              <div className="relative mb-4">
                <label
                  htmlFor="body"
                  className="tracking-midwest text-xs text-gray-400"
                >
                  BODY
                </label>
                <textarea
                  id="body"
                  name="body"
                  className="w-full bg-white rounded border border-gray-300 focus:border-red-500 focus:ring-2 focus:ring-red-200 h-32 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out"
                  onChange={(e) => setBody(e.target.value)}
                  value={body}
                />
              </div>
            )}
            {connected === true ? (
              <button
                onClick={() => sendRequest()}
                className="mr-2 text-white bg-red-500 border-0 py-2 px-6 focus:outline-none hover:bg-red-600 rounded text-base"
              >
                Send Request
              </button>
            ) : (
              <button
                onClick={() => connect()}
                className="mr-2 text-white bg-red-500 border-0 py-2 px-6 focus:outline-none hover:bg-red-600 rounded text-base"
              >
                Connect
              </button>
            )}
            {connected === true && (
              <button
                onClick={() => disconnect()}
                className="mr-2 text-white bg-red-500 border-0 py-2 px-6 focus:outline-none hover:bg-red-600 rounded text-base"
              >
                Disconnect
              </button>
            )}
            <button
              onClick={props.close}
              className="text-white bg-red-500 border-0 py-2 px-6 focus:outline-none hover:bg-red-600 rounded text-base"
            >
              Close
            </button>
          </div>
        </div>
      </section>
    );
  }
}

export default SendWebSocket;
