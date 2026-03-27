import React from 'react';

function TrackingDetails() {
  return (
    <div className='bg-black rounded-[10px] text-white p-4 flex flex-col gap-6'>
      <h2 className='text-violet-400 font-bold'>Track Your Package!</h2>
      <p className='text-gray-300'>
        Stay updated on your delivery's journey in real time. Enter your
        tracking ID to see the latest status and location!
      </p>
      <i className='text-cyan-200'>
        Server establishes a websocket connection with the courier and streams
        their location in real time, to the user's session.
      </i>
    </div>
  );
}

export default TrackingDetails;
