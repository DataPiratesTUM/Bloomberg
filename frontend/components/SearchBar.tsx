export function SearchBar() {
  return (
    <div className="flex items-center py-2.5 px-8 text-slate-400 rounded-lg bg-white">
      <svg
        className="mr-2 h-5 w-5 stroke-slate-500"
        fill="none"
        viewBox="0 0 24 24"
        strokeWidth="2"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        ></path>
      </svg>
      Search for theories...
    </div>
  );
}
