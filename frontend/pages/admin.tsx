import Head from "next/head";
import Image from "next/image";
import { Line } from "@nivo/line";
import { Layout } from "../components/Layout";

import React, { useEffect, useState } from "react";
import { Graph } from "../components/Graph";
import Link from "next/link";

import {
  useQuery,
  useQueryClient,
  useMutation,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { Match } from "../components/Match";

// export async function getServerSideProps() {
//   const res = await fetch(`http://localhost:3002/security/all`, {
//     headers: { "X-User-Id": "4e805cc9-fe3b-4649-96fc-f39634a557cd" },
//   });
//   const securities = res.json();
//   return { props: { securities } };
// }

export default function Admin() {
  const queryClient = useQueryClient();
  const [speed, setSpeed] = useState(1000);
  const [simulate, setSimulate] = useState(false);
  const varyingPrice = (x: number) =>
    0.35 +
    (1 / 8) *
      (Math.sin(2) - Math.sin(3 * x) + Math.sin(5 * x) - Math.sin(7 * x) + Math.sin(11 * x));

  const matches = useQuery(["matches"], async () => {
    const res = await fetch("http://localhost:3001/order/history", {
      headers: { "X-User-Id": "4e805cc9-fe3b-4649-96fc-f39634a557cd" },
    });
    const history: Match[] = await res.json();
    return history;
  });
  const securities = useQuery(["securities"], async () => {
    const res = await fetch("http://localhost:3002/security/all", {
      headers: { "X-User-Id": "4e805cc9-fe3b-4649-96fc-f39634a557cd" },
    });
    const s: Security[] = await res.json();
    return s;
  });

  const orderMutation = useMutation(
    (order: Order) =>
      fetch(`http://localhost:3001/order/place`, {
        method: "POST",
        body: JSON.stringify(order),
        headers: {
          "X-User-Id": "4e805cc9-fe3b-4649-96fc-f39634a557cd",
        },
      }),
    {
      onSuccess: () => queryClient.invalidateQueries(["matches"], {}, { cancelRefetch: false }),
    }
  );
  useEffect(() => {
    const interval = setInterval(() => {
      const order: Order = {
        // securities.data
        security:
          //   ? securities.data[Math.floor(Math.random() * securities.data.length)].security_id :
          "3e8b7701-9d3e-407a-b78a-d8fa4d07bff5",
        price: Math.floor((varyingPrice(Date.now()) + Math.random() * 0.1) * 1000),
        quantity: Math.floor(Math.random() * 1000),
        side: Math.random() < 0.5 ? "sell" : "buy",
      };
      simulate && orderMutation.mutate(order);
    }, speed);
    return () => clearInterval(interval);
  }, [simulate, speed, orderMutation, securities]);

  return (
    <>
      <Head>
        <title>Admin</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Layout>
        <h2 className="text-4xl font-bold tracking-tight  sm:text-6xl pb-4">Admin dashboard</h2>
        <label
          className="  px-3 
            py-4
          text-base
            font-normal
            inline-block
            text-gray-700"
          htmlFor="speed"
        >
          Trading Speed:
        </label>
        <input
          type="range"
          name="speed"
          id="speed"
          className="translate-y-1"
          value={10000 / speed}
          onChange={(e) => setSpeed(10000 / Number(e.target.value))}
          min={1}
          max={100}
        />
        {"  -->  "}
        {(100 / speed).toFixed(2)} transactions / second
        <br />
        <button
          className={` 
            px-3
            py-1.5
            text-base
            font-normal
            text-gray-700
            ${simulate ? "bg-red-400" : "bg-white"} bg-clip-padding
            border border-solid border-gray-300
            rounded
            transition
            ease-in-out
            m-0`}
          onClick={() => setSimulate((s) => !s)}
        >
          {simulate ? "Stop Autotrader" : "Start Autotrader"}
        </button>
        <br />
        <br />
        {matches.isLoading
          ? "Loading..."
          : matches.isError
          ? "Error!"
          : matches.data
          ? matches.data.map((match) => {
              return <Match key={match.created} match={match} />;
            })
          : null}
      </Layout>
    </>
  );
}