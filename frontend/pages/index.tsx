import Head from "next/head";
import Image from "next/image";
import { Line } from "@nivo/line";
import { Layout } from "../components/Layout";

import React, { useState } from "react";
import { Graph } from "../components/Graph";
import Link from "next/link";
import { TrendingList } from "../components/TrendingList";
import query from "../query";
import { useEffect } from "react";

export async function getServerSideProps() {
  let res_user = fetch(
    "https://organisation.ban.app/user/4e805cc9-fe3b-4649-96fc-f39634a557cd",
    {}
  );
  let res_user_securities = fetch("https://organisation.ban.app/security/all", {
    headers: {
      "X-User-Id": "4e805cc9-fe3b-4649-96fc-f39634a557cd",
    },
  });
  let res_user_portfolio = fetch("https://transaction.ban.app/order/value", {
    headers: {
      "X-User-Id": "4e805cc9-fe3b-4649-96fc-f39634a557cd",
    },
  });

  let res_user_trending = fetch("https://transaction.ban.app/trending");

  let [json_user, json_user_securities, json_user_portfolio, json_user_trending] =
    await Promise.all([res_user, res_user_securities, res_user_portfolio, res_user_trending]);
  const user: User = await json_user.json();
  const securities: Security[] = await json_user_securities.json();
  const portfolio: Portfolio[] = await json_user_portfolio.json();
  const trending: Trending = await json_user_trending.json();
  return { props: { user, trending, securities, portfolio } };
}

interface Home {
  user: User;
  securities: Security[];
  portfolio: Portfolio[];
  trending: Trending;
}

export default function Home({ user, securities, portfolio, trending }: Home) {
  const [portfolioValue, setPortfolioValue] = useState(0);

  useEffect(() => {
    setPortfolioValue(portfolio.length === 0 ? 0 : portfolio[portfolio.length - 1].value / 1000);
  }, [portfolio]);

  return (
    <>
      <Head>
        <title>ResearchX</title>
        <meta name="description" content="Research Trading Platform" />
        <link rel="icon" href="/favicon.svg" />
      </Head>
      <Layout>
        <section className="grid grid-rows-[0.5fr_2fr_2fr] grid-cols-2 gap-2">
          <section>
            <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl pb-4 basis-3/4">
              Hello, {user.Name}
            </h2>
          </section>

          <section className="row-span-2">
            <h3 className="text-4xl font-bold tracking-tight mb-8 sm:text-5xl">Your assets</h3>
            <div className="text-xl mb-6">Your securities are worth {portfolioValue}??? in total</div>
            {securities.map((security) => {
              // const percentageReturn = (security.price / security.price_bought - 1) * 100;
              return (
                <Link
                  href={"/securities/" + security.security_id}
                  key={security.security_id}
                  className="  border shadow rounded my-2 p-4 flex justify-between items-center"
                >
                  <div>
                    <h3 className="text-xl">{security.title}</h3>
                    <p className="text-gray-600 ">{security.description.slice(0, 100) + "..."}</p>
                  </div>
                  <div className="flex justify-center items-center text-xl ">
                    {((security.quantity * security.price) / 1000).toFixed(2)}???
                    {/* <div
                      className={`p-2 rounded ml-5 ${
                        percentageReturn > 0 ? "bg-green-300" : "bg-red-300"
                      }`}
                    >
                      {percentageReturn.toFixed(2)}%
                    </div> */}
                  </div>
                </Link>
              );
            })}
          </section>
          <section className=" flex flex-col justify-center">
            <Graph
              timeseries={portfolio
                .sort((a, b) => b.time - a.time)
                .map((p) => {
                  return { timestamp: p.time * 1000, price: p.value / 1000 };
                })}
            />
          </section>
          <section className="col-start-1 col-span-2">
            <TrendingList />
          </section>
        </section>
      </Layout>
    </>
  );
}
