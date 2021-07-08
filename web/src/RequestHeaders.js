function RequestHeaders(props) {
  return(
    <div className="p-4 md:w-1/2 w-full">
      <div className="bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">{pluralize(Object.keys(props.headers).length, "HEADER", "S")}</h2>
        {Object.keys(props.headers).map((key, i) => {
          return(
            <div key={i} className="flex border-t border-gray-200 py-2 text-xs">
              <span className="text-gray-500">{key}</span>
              <span className="ml-auto text-gray-900">{props.headers[key]}</span>
            </div>
          )
        })}
      </div>
    </div>
  )
}

const pluralize = (count, noun, suffix = 's') =>
  `${count} ${noun}${count !== 1 ? suffix : ''}`

export default RequestHeaders;
