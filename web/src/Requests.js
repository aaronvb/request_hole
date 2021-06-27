import {
  useQuery,
  gql
} from "@apollo/client"
import Request from './Request'

const ALL_REQUESTS = gql`
  query GetAllRequests {
    requests {
      fields {
        method
        url
        time
      }
      headers
      param_fields {
        form
        query
        json
        json_array
      }
    }
  }
`;

function AllRequests() {
  const { loading, error, data } = useQuery(ALL_REQUESTS);
  if (loading) return (
    <div>Loading...</div>
  );
  if (error) return (
    <div>Failed to load.</div>
  );

  return data.requests.map(({ fields, headers, param_fields }) => (
    <Request fields={fields} headers={headers} param_fields={param_fields}/>
  ));
}

function Requests() {
  return(
    <section className="text-gray-600 bg-gray-100 body-font h-full">
      <div className="container px-5 py-12 mx-auto">
        <AllRequests />
      </div>
    </section>
  )
}

export default Requests;