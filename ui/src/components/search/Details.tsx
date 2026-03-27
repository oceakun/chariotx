import { useState, useRef, useEffect } from 'react';
import { FiArrowRight } from 'react-icons/fi';
import { TbMapOff } from 'react-icons/tb';

const SearchDetails = ({
  onSearch,
}: {
  onSearch?: (query: string) => Promise<string | null>;
}) => {
  const [query, setQuery] = useState('');
  const [error, setError] = useState<string | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const errorTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  useEffect(() => {
    inputRef.current?.focus();
  }, []);

  function showError(msg: string) {
    if (errorTimerRef.current) clearTimeout(errorTimerRef.current);
    setError(msg);
    errorTimerRef.current = setTimeout(() => setError(null), 4000);
  }

  async function submit() {
    if (!onSearch || !query.trim()) return;
    setError(null);
    const err = await onSearch(query);
    if (err) showError(err);
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') submit();
    if (e.key === 'Escape') setQuery('');
  };

  return (
    <div className='bg-surface-selected rounded-[10px] text-text p-4 flex-col gap-6 flex items-start'>
      <h1 className='text-search font-bold'>Map the surf!</h1>
      <p className='text-text-muted'>
        Just look around for something interesting!
      </p>
      <div className='w-full mt-4 flex flex-col gap-2'>
        <div className='relative'>
          <input
            ref={inputRef}
            type='text'
            className='px-4 py-2 pr-10 rounded-lg border border-input-border bg-input-bg text-text shadow focus:outline-none focus:ring-2 focus:ring-input-focus transition w-full'
            placeholder='Search for a location or route...'
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={handleKeyDown}
          />
          <button
            onClick={submit}
            className='absolute right-3 top-1/2 -translate-y-1/2 text-text-muted hover:text-text cursor-pointer transition'
          >
            <FiArrowRight size={18} />
          </button>
        </div>

        {error && (
          <div className='flex items-center gap-2 text-sm text-text-muted bg-surface rounded-lg px-3 py-2 border border-input-border'>
            <TbMapOff size={16} className='text-search shrink-0' />
            <span>{error}</span>
          </div>
        )}
      </div>
    </div>
  );
};

export default SearchDetails;
