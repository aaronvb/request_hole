import React, { useEffect, useState } from 'react';
import RequestHeaders from './RequestHeaders';
import RequestParams from './RequestParams';

function Details(props) {
  const iconDown = (
    <svg xmlns="http://www.w3.org/2000/svg" className="cursor-pointer h-8 w-8 hover:text-black" viewBox="0 0 20 20" fill="currentColor">
      <path fillRule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clipRule="evenodd" />
    </svg>
  )

  const iconUp = (
    <svg xmlns="http://www.w3.org/2000/svg" className="cursor-pointer h-8 w-8 hover:text-black" viewBox="0 0 20 20" fill="currentColor">
      <path fillRule="evenodd" d="M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z" clipRule="evenodd" />
    </svg>
  )
  return(
    <div onClick={props.toggleDetails} className="flex ml-auto text-gray-500">
      {props.showDetails ? iconDown : iconUp}
    </div>
  );
}

function Request(props) {
  const time = formatTimeAgo(new Date(props.created_at));
  const [showDetails, setShowDetails] = useState(props.showAllDetails);

  useEffect(() => {
    setShowDetails(props.showAllDetails)
  }, [props.showAllDetails]) // Update this component show details if the parent show ALL details changes

  return(
    <div className="shadow bg-white rounded-md py-4 px-4 flex flex-wrap md:flex-nowrap mb-3 animate-slide-right">
      <div className="md:w-56 md:mb-0 mb-6 flex-shrink-0 flex flex-col">
        <span className="self-start inline-block py-1 px-2 rounded bg-indigo-50 text-indigo-500 text-s font-semibold tracking-widest">
          {props.fields.method}
        </span>
        <div className="mt-1 text-gray-400 text-sm">{time}</div>
      </div>
      <div className="md:flex-grow">
        <div className="flex w-full mx-auto">
          <div>
            <h2 className="tracking-midwest text-xs text-gray-400">URL</h2>
            <h2 className="font-medium text-gray-800 title-font mb-5 text-xl">{props.fields.url}</h2>
          </div>
          <Details id={props.id} showDetails={showDetails} toggleDetails={() => setShowDetails(!showDetails)}/>
        </div>
        {showDetails ?
        <section className="text-gray-600 body-font border-t-2 pt-3 border-gray-100">
          <div className="container py-2 mx-auto">
            <div className="flex flex-wrap -m-4">
              <RequestHeaders headers={props.headers} />
              <RequestParams params={props.param_fields} />
            </div>
          </div>
        </section>
        : <div></div>}
      </div>
    </div>
  )
}

// Calculate relative time
// https://blog.webdevsimplified.com/2020-07/relative-time-format/
//
const formatter = new Intl.RelativeTimeFormat(undefined, {
  numeric: 'auto'
})

const DIVISIONS = [
  { amount: 60, name: 'seconds' },
  { amount: 60, name: 'minutes' },
  { amount: 24, name: 'hours' },
  { amount: 7, name: 'days' },
  { amount: 4.34524, name: 'weeks' },
  { amount: 12, name: 'months' },
  { amount: Number.POSITIVE_INFINITY, name: 'years' }
]

function formatTimeAgo(date) {
  let duration = (date - new Date()) / 1000

  for (let i = 0; i <= DIVISIONS.length; i++) {
    const division = DIVISIONS[i]
    if (Math.abs(duration) < division.amount) {
      return formatter.format(Math.round(duration), division.name)
    }
    duration /= division.amount
  }
}

export default Request;
