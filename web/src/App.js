import Requests from "./Requests";
import SendRequest from "./SendRequest";
import SendWebSocket from "./SendWebSocket";
import Header from "./Header";
import { useQuery, gql } from "@apollo/client";
import { useState, useEffect } from "react";

const filters = [
  "GET",
  "POST",
  "PUT",
  "PATCH",
  "DELETE",
  "HEAD",
  "OPTIONS",
  "RECEIVE",
];

export const PROTOCOL = gql`
  query GetServerInfo {
    serverInfo {
      protocol
    }
  }
`;

function App() {
  const { data } = useQuery(PROTOCOL);
  const [sendRequestVisible, setSendRequestVisible] = useState(false);
  const [protocol, setProtocol] = useState("");

  useEffect(() => {
    if (data) {
      setProtocol(data.serverInfo.protocol);
    }
  }, [data]);

  return (
    <div>
      <Header
        sendRequestVisible={sendRequestVisible}
        setSendRequestVisible={setSendRequestVisible}
      />
      {protocol === "ws" ? (
        <SendWebSocket
          visible={sendRequestVisible}
          close={() => setSendRequestVisible(false)}
        />
      ) : (
        <SendRequest
          filters={filters}
          visible={sendRequestVisible}
          close={() => setSendRequestVisible(false)}
        />
      )}

      <Requests filters={filters} />
    </div>
  );
}

export default App;
