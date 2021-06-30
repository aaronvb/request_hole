import {
  useQuery,
  useMutation,
  gql
} from "@apollo/client";
import Request from './Request';
import React, { useState, useEffect } from 'react';

const ALL_REQUESTS = gql`
  query GetAllRequests {
    requests {
      id
      fields {
        method
        url
      }
      headers
      param_fields {
        form
        query
        json
        json_array
      }
      created_at
    }
  }
`;

const REQUESTS_SUBSCRIPTION = gql`
  subscription OnRequestCreated {
    request {
      id
      fields {
        method
        url
      }
      headers
      param_fields {
        form
        query
        json
        json_array
      }
      created_at
    }
  }
`;

const CLEAR_REQUESTS = gql`
  mutation ClearRequests {
    clearRequests
  }
`;

function filterRequests(requests, filter = "All") {
  return requests.filter(request => !(filter !== "ALL" && filter !== request.fields.method))
}

function AllRequests(props) {
  if (props.loading) return (
    <div>Loading requests...</div>
  );

  if (props.error) return (
    <div>Failed to load.</div>
  );

  const sortedRequests = props.requests.slice().sort((a, b) =>
    new Date(b.created_at) - new Date(a.created_at)
  );

  return filterRequests(sortedRequests, props.selectedFilter).map(({id, fields, headers, param_fields, created_at}) => (
    <Request
      key={id}
      created_at={created_at}
      fields={fields}
      headers={headers}
      param_fields={param_fields}
      id={id}
      showAllDetails={props.showAllDetails}
    />
  ));
}

function ToggleDetails(props) {
  const iconHide = (
    <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
    </svg>
  )

  const iconShow = (
    <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
    </svg>
  )

  return(
    <div onClick={props.toggle} className="ml-1 items-center cursor-pointer inline-flex bg-indigo-500 border-0 py-1 px-3 focus:outline-none hover:bg-indigo-900 rounded text-white">
      {props.showAllDetails ? iconHide : iconShow}
      {props.showAllDetails ? "Hide Details" : "Show Details"}
    </div>
  )
}

function Filters(props) {
  return props.filters.map((filter, i) => (
    <li onClick={() => props.setSelectedFilter(filter)}>
      <button class={`${i === (props.filters.length - 1) ? "rounded-b" : ""} bg-indigo-500 hover:bg-indigo-900 py-2 px-4 block w-full text-left whitespace-no-wrap`}>
        {filter}
      </button>
    </li>
  ));
}

function Requests() {
  const {loading, error, data, subscribeToMore} = useQuery(ALL_REQUESTS);
  const [clearRequests] = useMutation(CLEAR_REQUESTS, {
    update(cache) {
      cache.modify({
        fields: {
          requests() {
            return []
          }
        }
      })
    }
  });

  const [requests, setRequests] = useState([]);
  const [subscribed, setSubscribed] = useState(false);
  const [showAllDetails, setShowAllDetails] = useState(true);
  const [selectedFilter, setSelectedFilter] = useState("ALL");
  const filters = [
    "GET",
    "POST",
    "PUT",
    "PATCH",
    "DELETE",
    "COPY",
    "HEAD",
    "OPTIONS",
    "LINK",
    "UNLINK",
    "PURGE",
    "LOCK",
    "UNLOCK",
    "PROPFIND",
    "VIEW",
  ]

  useEffect(() => {
    if (data) {
      setRequests(data.requests);
    }

    if (!subscribed) {
      subscribeToMore({
        document: REQUESTS_SUBSCRIPTION,
        updateQuery: (prev, { subscriptionData }) => {
          if (!subscriptionData.data) return prev
          const newRequest = subscriptionData.data.request
          return Object.assign({}, prev, {
                requests: [newRequest, ...prev.requests]
            });
        }
      });
      setSubscribed(true);
    }
  }, [data, subscribed, subscribeToMore]);

  return(
    <section className="text-gray-600 bg-gray-100 body-font h-full">
      <div className="container px-5 py-12 mx-auto">
        <div className="flex flex-wrap w-full">
          <div className="lg:w-1/2 w-full mb-6 lg:mb-0">
            <div className="flex flex-col sm:flex-row sm:items-center items-start mx-auto">
              <h1 className="sm:text-2xl text-xl font-medium title-font mb-2 text-gray-900">
                {pluralize(filterRequests(requests, selectedFilter).length, "Request")}
              </h1>
            </div>
            <div className="h-1 w-1/6 bg-indigo-500 rounded mb-4"></div>
          </div>
          <div className="flex items-center lg:w-1/2 w-full mb-5 flex-row-reverse">
            <div className="group inline-block relative">
              <button className="ml-1 items-center cursor-pointer inline-flex bg-indigo-500 border-0 py-1 px-3 focus:outline-none hover:bg-indigo-900 rounded text-white">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
                </svg>
                Filter: {selectedFilter}
              </button>
              <ul class="absolute hidden right-0 w-max text-white pt-1 group-hover:block z-10">
                <li onClick={() => setSelectedFilter("ALL")}>
                  <button class="rounded-t bg-indigo-500 hover:bg-indigo-900 py-2 px-4 block w-full text-left whitespace-no-wrap">
                    ALL
                  </button>
                </li>
                <Filters filters={filters} setSelectedFilter={setSelectedFilter} />
              </ul>
            </div>
            <ToggleDetails showAllDetails={showAllDetails} toggle={() => setShowAllDetails(!showAllDetails)}/>
            <div onClick={() => {if (window.confirm('Are you sure you want to clear all requests?')) clearRequests()}} className="cursor-pointer items-center inline-flex bg-indigo-500 border-0 py-1 px-3 focus:outline-none hover:bg-indigo-900 rounded text-white">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
              Clear Requests
            </div>
          </div>
        </div>
        <AllRequests selectedFilter={selectedFilter} error={error} loading={loading} requests={requests} showAllDetails={showAllDetails} />
      </div>
    </section>
  )
}

const pluralize = (count, noun, suffix = 's') =>
  `${count} ${noun}${count !== 1 ? suffix : ''}`;

export default Requests;