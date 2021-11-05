import Requests from "./Requests";
import SendRequest from "./SendRequest";
import Header from "./Header";
import { useState } from "react";

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

function App() {
  const [sendRequestVisible, setSendRequestVisible] = useState(false);
  return (
    <div>
      <Header
        sendRequestVisible={sendRequestVisible}
        setSendRequestVisible={setSendRequestVisible}
      />
      <SendRequest
        filters={filters}
        visible={sendRequestVisible}
        close={() => setSendRequestVisible(false)}
      />
      <Requests filters={filters} />
    </div>
  );
}

export default App;
