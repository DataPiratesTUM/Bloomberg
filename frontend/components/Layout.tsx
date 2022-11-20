import Image from "next/image";
import Link from "next/link";
import React, { ChangeEventHandler, ReactNode, useState } from "react";
import Logo from "../assets/Logo";
import Profile from "../assets/Profile";
import { SearchBar } from "./SearchBar";
interface Result {
  Id: string;
  Name: string;
}
export function Layout({ children }: { children: ReactNode }) {
  const [isSearching, setSearching] = useState(false);
  const [searchText, setSearchText] = useState("");
  const [response, setResponse] = useState<Result[] | null>(null);
  const searchFunc = async (data: string) => {
    var requestOptions = {
      method: "GET",
    };

    const dataReq = await fetch(
      "http://localhost:3002/security/search/title?query=" + data,
      requestOptions
    )
      .then((response) => response.json())
      .then((result) => setResponse(result))
      .catch((error) => console.log("error", error));
    return dataReq;
  };

  const inputHandler = (e: any) => {
    var search = e.target.value;
    setSearchText(search);
    search.length >= 1 ? setSearching(true) : setSearching(false);
    searchFunc(search);
  };
  console.log("RESPONSE COOOOLEEE" + response);

  const pageSearching = (
    <div>
      {response?.length === 0 || response == null ? (
        <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl pb-4 basis-3/4 text-center m-20">
          NOTHING FOUND
        </h2>
      ) : (
        <>
          <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl pb-4 basis-3/4">
            Search Results
          </h2>
          {response!.map((result) => {
            return (
              <Link href={"/securities/" + result.Id} key={result.Id}>
                <div className=" m-2 border shadow rounded my-2 p-4 flex justify-between">
                  <p>{result.Name}</p>
                </div>
              </Link>
            );
          })}
        </>
      )}
    </div>
  );

  const isLoadingPage = isSearching ? pageSearching : children;

  return (
    <>
      <header className="grid grid-cols-3 p-6 bg-slate-800 text-white items-center ">
        <Link href="/" className="flex">
          <Logo />
          <h1 className="px-6 text-xl ">ResearchX</h1>
        </Link>
        <section className="flex justify-center">
          <SearchBar inputHandler={inputHandler} />
        </section>

        <section className="flex justify-end">
          <Profile />
        </section>
      </header>

      <main className="p-8">{isLoadingPage}</main>
    </>
  );
}
