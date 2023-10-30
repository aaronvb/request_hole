import ReactJson from "react-json-view";

function RequestParams(props) {
  if (props.params && props.params.json) {
    return <JsonParams json={props.params.json} />;
  } else if (props.params && props.params.json_array) {
    return <JsonParams json={props.params.json_array} />;
  } else if (props.params && props.params.query) {
    return <QueryParams query={props.params.query} />;
  } else if (props.params && props.params.form) {
    return <FormParams form={props.params.form} />;
  } else if (props.message) {
    return <Message body={props.message} />;
  } else {
    return (
      <div className="p-4 md:w-1/2 w-full">
        <div className="bg-gray-100 p-4 rounded">
          <h2 className="tracking-midwest text-xs text-gray-400">NO PARAMS</h2>
        </div>
      </div>
    );
  }
}

function QueryParams(props) {
  return (
    <div className="p-4 md:w-1/2 w-full">
      <div className="bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">
          {pluralize(Object.keys(props.query).length, "QUERY PARAM", "S")}{" "}
        </h2>
        {Object.keys(props.query).map((key, i) => {
          return (
            <div key={i} className="flex border-t border-gray-200 py-2 text-xs">
              <span className="text-gray-500">{key}</span>
              <span className="ml-auto text-gray-900">{props.query[key]}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
}

function FormParams(props) {
  return (
    <div className="p-4 md:w-1/2 w-full">
      <div className="bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">
          {pluralize(Object.keys(props.form).length, "FORM PARAM", "S")}
        </h2>
        {Object.keys(props.form).map((key, i) => {
          return (
            <div key={i} className="flex border-t border-gray-200 py-2 text-xs">
              <span className="text-gray-500">{key}</span>
              <span className="ml-auto text-gray-900">{props.form[key]}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
}

function JsonParams(props) {
  return (
    <div className="p-4 md:w-1/2 w-full">
      <div className="h-full bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">
          JSON BODY
        </h2>
        <div className="flex border-t border-gray-200 py-2 text-xs">
          <ReactJson src={props.json} name={false} />
        </div>
      </div>
    </div>
  );
}

function Message(props) {
  return (
    <div className="p-4 md:w-1/2 w-full">
      <div className="h-full bg-gray-100 p-4 rounded">
        <h2 className="tracking-midwest text-xs text-gray-400 mb-2">MESSAGE</h2>
        <div className="flex border-t border-gray-200 py-2 text-xs">
          {renderJSONOrString(props.body)}
        </div>
      </div>
    </div>
  );
}

const pluralize = (count, noun, suffix = "s") =>
  `${count} ${noun}${count !== 1 ? suffix : ""}`;

function renderJSONOrString(message) {
  try {
    const json = JSON.parse(message);
    return <ReactJson src={json} name={false} />;
  } catch (e) {
    return message;
  }
}

export default RequestParams;
