import { useQuery, gql } from "@apollo/client";
import { useState, useEffect } from "react";

export const SERVER_INFO = gql`
  query GetServerInfo {
    serverInfo {
      request_address
      request_port
      build_info
      protocol
    }
  }
`;

function ServerInfo(props) {
  if (props.loading) return <div>Loading server info...</div>;

  if (props.error) return <div>Failed to load server info.</div>;

  return (
    <div className="bg-gray-100 rounded py-1 px-3 text-sm flex flex-wrap items-center justify-center">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-4 w-4 mr-1"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth={2}
          d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"
        />
      </svg>
      Listening on: {props.url}
    </div>
  );
}

function Header(props) {
  const { loading, error, data } = useQuery(SERVER_INFO);
  const [url, setUrl] = useState("");
  const [version, setVersion] = useState("");
  const [protocol, setProtocol] = useState("");

  useEffect(() => {
    if (data) {
      setUrl(
        `${data.serverInfo.protocol}://${data.serverInfo.request_address}:${data.serverInfo.request_port}`
      );
      setVersion(data.serverInfo.build_info["version"]);
      setProtocol(data.serverInfo.protocol);
    }
  }, [data]);

  return (
    <header className="text-gray-600 body-font border-b-2 bg-white">
      <div className="container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center">
        <a
          href="/"
          className="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0"
        >
          <span className="text-xl">Request Hole</span>
          <h2 className="tracking-widest text-sm ml-2 title-font font-light text-gray-400">
            {version}
          </h2>
        </a>
        <div className="md:mr-auto md:ml-4 md:py-1 md:pl-4 md:border-l md:border-gray-400	flex flex-wrap items-center text-base justify-center">
          <ServerInfo loading={loading} error={error} url={url} />
        </div>
        <nav className="md:ml-auto flex flex-wrap items-center text-base justify-center">
          <button
            onClick={() =>
              props.setSendRequestVisible(!props.sendRequestVisible)
            }
            className="focus:outline-none mr-5 hover:text-gray-900 flex flex-wrap items-center text-base"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5 mr-1"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path d="M8.707 7.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l2-2a1 1 0 00-1.414-1.414L11 7.586V3a1 1 0 10-2 0v4.586l-.293-.293z" />
              <path d="M3 5a2 2 0 012-2h1a1 1 0 010 2H5v7h2l1 2h4l1-2h2V5h-1a1 1 0 110-2h1a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V5z" />
            </svg>
            Send a {protocol === "ws" ? "WebSocket Message" : "Request"}
          </button>
          <a
            href="https://github.com/aaronvb/request_hole"
            className="hover:text-gray-900 flex flex-wrap items-center text-base"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5 mr-1"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fillRule="evenodd"
                d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z"
                clipRule="evenodd"
              />
            </svg>
            View Project on GitHub
          </a>
        </nav>
      </div>
    </header>
  );
}

export default Header;
