import { useState, useRef, useEffect } from 'react';

const SearchDetails = ({
  onSearch,
}: {
  onSearch?: (query: string) => void;
}) => {
  const [query, setQuery] = useState('');
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (inputRef.current) {
      inputRef.current.focus();
    }
  }, []);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && onSearch) {
      onSearch(query);
    }
    if (e.key === 'Escape') {
      setQuery('');
    }
  };

  return (
    <div className='bg-black rounded-[10px] text-white p-4 flex-col gap-6 flex items-start'>
      <h1 className='text-yellow-400 font-bold'>Map the surf!</h1>
      <p className='text-gray-300'>
        Just look around for something interesting!
      </p>
      <div className='w-full mt-4'>
        <input
          ref={inputRef}
          type='text'
          className='px-4 py-2 rounded-full border border-gray-300 shadow focus:outline-none focus:ring-2 focus:ring-blue-400 transition w-full'
          placeholder='Search for a location or route...'
          value={query}
          onChange={handleInputChange}
          onKeyDown={handleKeyDown}
        />
      </div>
    </div>
  );
};

export default SearchDetails;
