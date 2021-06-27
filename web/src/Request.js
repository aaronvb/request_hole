import ReactJson from 'react-json-view'

function Headers(props) {
  return(
    <div class="p-4 md:w-1/2 w-full">
      <div class="bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">{Object.keys(props.headers).length} HEADERS</h2>
        {Object.keys(props.headers).map((key, i) => {
          return(
            <div key={i} class="flex border-t border-gray-200 py-2 text-xs">
              <span class="text-gray-500">{key}</span>
              <span class="ml-auto text-gray-900">{props.headers[key]}</span>
            </div>
          )
        })}
      </div>
    </div>
  )
}

function Params(props) {
  if (props.params.json) {
    return <JsonParams json={props.params.json} />
  } else if (props.params.json_array) {
    return <JsonParams json={props.params.json_array} />
  } else if (props.params.query) {
    return <QueryParams query={props.params.query} />
  } else if (props.params.form) {
    return <FormParams form={props.params.form} />
  } else {
    return (
      <div class="p-4 md:w-1/2 w-full">
        <div class="bg-gray-100 p-4 rounded">
          <h2 className="tracking-midwest text-xs text-gray-400">NO PARAMS</h2>
        </div>
      </div>
    )
  }
}

function QueryParams(props) {
  return (
    <div class="p-4 md:w-1/2 w-full">
      <div class="bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">{Object.keys(props.query).length} QUERY PARAMS</h2>
          {Object.keys(props.query).map((key, i) => {
            return(
              <div key={i} class="flex border-t border-gray-200 py-2 text-xs">
                <span class="text-gray-500">{key}</span>
                <span class="ml-auto text-gray-900">{props.query[key]}</span>
              </div>
            )
          })}
      </div>
    </div>
  )
}

function FormParams(props) {
  return (
    <div class="p-4 md:w-1/2 w-full">
      <div class="bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">{Object.keys(props.form).length} FORM PARAMS</h2>
          {Object.keys(props.form).map((key, i) => {
            return(
              <div key={i} class="flex border-t border-gray-200 py-2 text-xs">
                <span class="text-gray-500">{key}</span>
                <span class="ml-auto text-gray-900">{props.form[key]}</span>
              </div>
            )
          })}
      </div>
    </div>
  )
}

function JsonParams(props) {
  return (
    <div class="p-4 md:w-1/2 w-full">
      <div class="h-full bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">JSON BODY</h2>
        <div class="flex border-t border-gray-200 py-2 text-xs">
          <ReactJson src={props.json} name={false}/>
        </div>
      </div>
    </div>
  )
}

function Request(props) {
  const time = new Date(props.fields.time).toLocaleString('en-US', {hour12: false}).replace(', ', ' - ')
  return(
    <div className="shadow bg-white rounded-md py-4 px-4 flex flex-wrap md:flex-nowrap mb-3 animate-fade-in-down">
      <div className="md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col">
        <span className="self-start inline-block py-1 px-2 rounded bg-indigo-50 text-indigo-500 text-s font-semibold tracking-widest">
          {props.fields.method}
        </span>
        <span className="mt-1 text-gray-400 text-sm">{time}</span>
      </div>
      <div className="md:flex-grow">
        <div className="border-b-2 mb-3 border-gray-100">
          <h2 className="tracking-midwest text-xs text-gray-400">URL</h2>
          <h2 className="font-medium text-gray-800 title-font mb-5 text-xl">{props.fields.url}</h2>
        </div>
        <section class="text-gray-600 body-font">
          <div class="container py-2 mx-auto">
            <div class="flex flex-wrap -m-4">
              <Headers headers={props.headers} />
              <Params params={props.param_fields} />
            </div>
          </div>
        </section>
      </div>
    </div>
  )
}

export default Request;
