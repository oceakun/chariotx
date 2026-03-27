import React from 'react';

function TrackingDetails() {
  return (
    <div className='bg-surface-selected rounded-[10px] text-text p-4 flex flex-col gap-6'>
      <h2 className='text-tracking font-bold'>Track Your Package!</h2>
      <p className='text-text-muted'>
        Stay updated on your delivery&apos;s journey in real time. Enter your
        tracking ID to see the latest status and location!
      </p>
      <i className='text-tracking-note'>
        Server establishes a websocket connection with the courier and streams
        their location in real time, to the user&apos;s session.
      </i>
    </div>
  );
}

export default TrackingDetails;
