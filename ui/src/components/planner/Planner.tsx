import { FaRoute } from 'react-icons/fa';

const Planner = ({
  onClick,
  selected,
}: {
  onClick?: () => void;
  selected?: boolean;
}) => (
  <button
    aria-label='Planner'
    className={`p-2 rounded-full transition hover:bg-black cursor-pointer flex items-center justify-center${selected ? ' bg-black' : ''}`}
    onClick={onClick}
  >
    <FaRoute
      size={24}
      className={`${selected ? 'text-cyan-400' : 'text-cyan-200'} hover:text-cyan-400`}
    />
  </button>
);

export default Planner;
