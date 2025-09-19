import React from 'react'

function LoadingScreen() {
  return (
    <div className="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-green-700 via-blue-400 to-yellow-200 opacity-60 z-20">
      <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-white"></div>
    </div>
  )
}

export default LoadingScreen