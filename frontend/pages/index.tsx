import Head from "next/head";
import Image from "next/image";
import { Line } from "@nivo/line";
import { Layout } from "../components/Layout";

import React, { useState } from "react";
import { Graph } from "../components/Graph";
import Link from "next/link";
import { TrendingList } from "../components/TrendingList";
import query from "../query";

export async function getServerSideProps() {
  // const res = await fetch("");
  //let user: User = await res.json();

  // Mock data
  let user: User = {
    user_id: "123",
    name: "Cole Friedlaender",
    balance: 12345,
    securities: [
      {
        name: "Does the Higgs-Boson exist? ",
        id: "1234",
        qty: 4,
        price: 938,
        price_bought: 378,
      },
      {
        name: "Does sitting during a hackathon for 24 hours cause diabetis?",
        id: "1235",
        qty: 38,
        price: 333,
        price_bought: 546,
      },
      {
        name: "Does sitting during a hackathon for 24 hours cause diabetis?",
        id: "1236",
        qty: 38,
        price: 333,
        price_bought: 546,
      },
      {
        name: "Does sitting during a hackathon for 24 hours cause diabetis?",
        id: "1237",
        qty: 38,
        price: 333,
        price_bought: 546,
      },
      {
        name: "Does sitting during a hackathon for 24 hours cause diabetis?",
        id: "1238",
        qty: 38,
        price: 333,
        price_bought: 546,
      },
    ],
    timeseries: [
      { timestamp: Date.now(), price: 567 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 1, price: 670 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 2, price: 589 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 3, price: 400 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 4, price: 567 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 5, price: 670 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 6, price: 589 },
      { timestamp: Date.now() + 1000 * 60 * 60 * 24 * 7, price: 400 },
    ],
  };

  // Mock data
  let trending: TrendingList = {
    trendings: [
      {
        security_id: "1",
        title: "Test1",
      },
      {
        security_id: "2",
        title: "Test2",
      },
      {
        security_id: "3",
        title: "Test3",
      },
      {
        security_id: "4",
        title: "Test4",
      },
    ],
  };

  return { props: { user, trending } };
}

interface Home {
  user: User;
  trending: TrendingList;
}
interface Result {
  Id: string;
  Name: string;
}

export default function Home(props: Home) {
  const { user } = props;
  const { trending } = props;
  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: "top" as const,
      },
      title: {
        display: true,
        text: "Chart.js Line Chart",
      },
    },
  };

  const pageNotSearching = (
    <section className="grid grid-rows-[0.5fr_2fr_2fr] grid-cols-2 gap-2">
      <section>
        <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl pb-4 basis-3/4">
          Hello, {user.name}{" "}
        </h2>
      </section>

      <section className="row-span-2">
        <h3 className="text-4xl font-bold tracking-tight  sm:text-5xl">
          Your assets
        </h3>
        {user.securities.map((security) => {
          const percentageReturn =
            (security.price / security.price_bought - 1) * 100;
          return (
            <Link
              href={"/securities/" + security.id}
              key={security.id}
              className="  border shadow rounded my-2 p-4 flex justify-between items-center"
            >
              <p>{security.name}</p>
              <div className="flex justify-center items-center ">
                {(security.qty * security.price) / 100}€
                <div
                  className={`p-2 rounded ml-5 ${
                    percentageReturn > 0 ? "bg-green-300" : "bg-red-300"
                  }`}
                >
                  {percentageReturn.toFixed(2)}%
                </div>
              </div>
            </Link>
          );
        })}
      </section>
      <section className=" flex flex-col justify-center">
        <p className="text-xl mb-6">
          Your futures are worth {user.balance / 100}€ in total
        </p>

        <Graph timeseries={user.timeseries} />
      </section>
      <section className="col-start-1 col-span-2">
        <TrendingList trendingList={trending} />
      </section>
    </section>
  );

  return (
    <>
      <Head>
        <title>ResearchX</title>
        <meta name="description" content="Research Trading Platform" />
        <link rel="icon" href="/favicon.svg" />
      </Head>
      <Layout>{pageNotSearching}</Layout>
    </>
  );
}
