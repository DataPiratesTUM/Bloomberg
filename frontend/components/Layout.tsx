import Image from "next/image";
import Link from "next/link";
import React, { ReactNode } from "react";
import Logo from "../assets/Logo";
import Profile from "../assets/Profile";
import { SearchBar } from "./SearchBar";

export function Layout({ children }: { children: ReactNode }) {
  return (
    <>
      <header className="grid grid-cols-3 p-6 bg-slate-800 text-white items-center ">
        <Link href="/" className="flex">
          <Logo />
          <h1 className="px-6 text-xl ">ResearchX</h1>
        </Link>
        <section className="flex justify-center">
          <SearchBar />
        </section>

        <section className="flex justify-end">
          <Profile />
        </section>
      </header>

      <main className="p-8">{children}</main>
    </>
  );
}
