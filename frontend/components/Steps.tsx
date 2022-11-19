export default function Setps() {
  let test = 1;
  let test1 = 2;
  return (
    <div>
      <h2 className="sr-only">Steps</h2>

      <div className="relative after:absolute after:inset-x-0 after:top-1/2 after:block after:h-0.5 after:-translate-y-1/2 after:rounded-lg after:bg-gray-100">
        <ol className="relative z-10 flex justify-between text-sm font-medium text-gray-500">
          <li className="flex items-center bg-white p-2">
            <span className="h-6 w-6 rounded-full bg-gray-100 text-center text-[10px] font-bold leading-6">
              1
            </span>

            <span className="hidden sm:ml-2 sm:block"> Funding Phase </span>
          </li>

          <li className="flex items-center bg-white p-2">
            <span
              className={`${
                test + test1 == 2
                  ? "h-6 w-6 rounded-full bg-blue-600 text-center text-[10px] font-bold leading-6 text-white"
                  : "h-6 w-6 rounded-full bg-gray-100 text-center text-[10px] font-bold leading-6"
              }  `}
            >
              2
            </span>

            <span className="hidden sm:ml-2 sm:block"> Research Phase </span>
          </li>

          <li className="flex items-center bg-white p-2">
            <span className="h-6 w-6 rounded-full bg-gray-100 text-center text-[10px] font-bold leading-6">
              3
            </span>

            <span className="hidden sm:ml-2 sm:block"> Evaluation Phase </span>
          </li>
        </ol>
      </div>
    </div>
  );
}
