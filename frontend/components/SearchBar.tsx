/* export function SearchBar() {
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
 */

import { ChangeEventHandler, useState } from "react";

export function SearchBar({
  inputHandler,
}: {
  inputHandler: ChangeEventHandler;
}) {
  return (
    <div className="">
      <div className="w-4/4">
        <input
          type="text"
          className="
        form-control
        block
        w-full
        px-3
        py-1.5
        text-base
        font-normal
        text-gray-700
        bg-white bg-clip-padding
        border border-solid border-gray-300
        rounded
        transition
        ease-in-out
        m-0
        focus:text-gray-700 focus:bg-white focus:outline-none
      "
          id="exampleFormControlInput1"
          placeholder="Search"
          onChange={inputHandler}
        />
      </div>
    </div>
  );
}
