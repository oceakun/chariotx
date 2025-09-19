import { FaRegClock } from "react-icons/fa";

const Estimate = ({ onClick, selected }: { onClick?: () => void; selected?: boolean }) => (
  <button
    aria-label="Estimate"
    className={`p-2 rounded-full transition hover:bg-black cursor-pointer flex items-center justify-center${selected ? ' bg-black' : ''}`}
    onClick={onClick}
  >
    <FaRegClock size={24} className={`${selected ? 'text-orange-400' : 'text-orange-200'} hover:text-orange-400`} />
  </button>
);

export default Estimate;