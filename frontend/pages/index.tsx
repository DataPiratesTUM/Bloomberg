import Head from "next/head";
import Image from "next/image";
import { Line } from "@nivo/line";
import { Layout } from "../components/Layout";

import React from "react";
import { Graph } from "../components/Graph";

export async function getServerSideProps() {
  // const res = await fetch("");
  //let user: User = await res.json();

  // Mock data
  let user: User = {
    id: "123",
    name: "Cole Friedlaender",
    balance: 12345,
    securities: [
      { name: "Does the Higgs-Boson exist? ", id: "1234", qty: 4, price: 938, price_bought: 378 },
      {
        name: "Does sitting during a hackathon for 24 hours cause diabetis?",
        id: "1234",
        qty: 38,
        price: 333,
        price_bought: 546,
      },
    ],
    timeseries: [
      { timestamp: 1668821699865, price: 567 },
      { timestamp: 1668821799865, price: 670 },
      { timestamp: 1668821899865, price: 589 },
      { timestamp: 1668821999865, price: 400 },
    ],
  };

  return { props: { user } };
}

export default function Home(props: Home) {
  const { user } = props;

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

  return (
    <>
      <Head>
        <title>To be determined</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Layout>
        <h2 className="text-4xl font-bold tracking-tight  sm:text-6xl pb-4">Hello, {user.name}</h2>

        <p className="text-xl">Your futures are worth {user.balance / 100}€ in total</p>
        <Graph timeseries={user.timeseries} />

        <h2 className="text-4xl font-bold tracking-tight  sm:text-6xl pb-4">Your assets</h2>
        {user.securities.map((security) => (
          <section key={security.id} className="border shadow rounded my-2 p-4 ">
            {security.name} {(security.qty * security.price) / 100}€{" "}
            {(security.price / security.price_bought).toFixed(2)}%
          </section>
        ))}
      </Layout>
    </>
  );
}
