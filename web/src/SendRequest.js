import { useState, useEffect } from 'react';

function Filters(props) {
  return props.filters.map((filter, i) => (
    <option key={i}>{filter}</option>
  ));
}

function SendRequest(props) {
  const [ method, setMethod ] = useState("GET");
  const [ url, setUrl ] = useState("");
  const [ body, setBody ] = useState(JSON.stringify({"hello": "world"}));

  const sendRequest = () => {
    fetch(url,
      {
        method: method,
        body: (method === "GET" || method === "HEAD") ? null : body,
        headers: {
          'Content-Type': 'application/json'
        }
      }
    )
  };

  useEffect(() => {
    setUrl(`http://${props.address}:${props.port}`);
  }, [props.address, props.port]);

  if (!props.visible) {
    return (
      <div></div>
    )
  } else {
    return(
      <section className="text-gray-600 bg-gray-100 body-font h-full">
        <div className="container p-5 mx-auto max-w-2xl">
          <div className="bg-white rounded shadow py-4 px-4">
            <h2 className="text-gray-900 text-lg mb-1 font-medium title-font">Send a Request</h2>
            <div className="flex flex-wrap mb-4">
              <div className="md:pr-1 md:w-2/6 sm:w-1/2 w-full">
                <label htmlFor="method" className="tracking-midwest text-xs text-gray-400">METHOD</label>
                <div className="flex">
                  <div className="relative w-full">
                    <select
                      name="method"
                      id="method"
                      className="w-full rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-red-200 focus:border-red-500 text-base pl-3 pr-10"
                      onChange={(e) => setMethod(e.target.value)}
                      value={method}>
                    <Filters filters={props.filters} />
                    </select>
                    <span className="absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center">
                      <svg fill="none" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" className="w-4 h-4" viewBox="0 0 24 24">
                        <path d="M6 9l6 6 6-6"></path>
                      </svg>
                    </span>
                  </div>
                </div>
              </div>
              <div className="md:pl-1 md:w-4/6 sm:w-1/2 w-full">
                <div className="relative">
                  <label htmlFor="url" className="tracking-midwest text-xs text-gray-400">URL</label>
                  <input
                    type="text"
                    id="url"
                    name="url"
                    className="w-full rounded border border-gray-300 focus:border-red-500 focus:bg-white focus:ring-2 focus:ring-red-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                  />
                </div>
              </div>
            </div>
            <div className="relative mb-4">
              <label htmlFor="body" className="tracking-midwest text-xs text-gray-400">BODY</label>
              <textarea
                id="body"
                name="body"
                className="w-full bg-white rounded border border-gray-300 focus:border-red-500 focus:ring-2 focus:ring-red-200 h-32 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out"
                onChange={(e) => setBody(e.target.value)}
                value={body} />
            </div>
            <button onClick={() => sendRequest()} className="mr-2 text-white bg-red-500 border-0 py-2 px-6 focus:outline-none hover:bg-red-600 rounded text-base">Send Request</button>
            <button onClick={props.close} className="text-white bg-red-500 border-0 py-2 px-6 focus:outline-none hover:bg-red-600 rounded text-base">Close</button>
          </div>
        </div>
      </section>
    )
  }
}

export default SendRequest;
