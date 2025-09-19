import { FaCarSide } from "react-icons/fa";

const Cab = ({ onClick, selected }: { onClick?: () => void; selected?: boolean }) => (
  <button
    aria-label="Cab"
    className={`p-2 rounded-full transition hover:bg-black hover:cursor-pointer flex items-center justify-center${selected ? ' bg-black' : ''}`}
    onClick={onClick}
  >
    <FaCarSide size={24} className={`${selected ? 'text-green-400' : 'text-green-200'} hover:text-green-400`} />
  </button>
);

export default Cab;
