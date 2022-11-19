import { formatDistance, formatDistanceToNow } from "date-fns";
import { GetServerSidePropsContext, NextPageContext } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import { useState } from "react";
import { Graph } from "../../components/Graph";
import { Layout } from "../../components/Layout";
import Setps from "../../components/Steps";
import Link from "next/link";

export function getServerSideProps(context: GetServerSidePropsContext) {
  const security_id = context.params?.id;

  // const res = await fetch("");
  //let user: User = await res.json();

  // Mock data
  let security: Security = {
    security_id: "12345",
    creation_date: Date.now() - 100000,
    price: 124,
    creator: {
      name: "Technical University Munich",
      organisation_id: "TUM1234",
    },
    description: "Best University in the World",
    title: "Does the Higgs-Boson exist?",
    orders: [
      {
        id: "edidjw",
        price: 234,
        qty: 3,
        side: "SELL",
      },
      {
        id: "edifjw",
        price: 23,
        qty: 245,
        side: "SELL",
      },
      {
        id: "edrdjw",
        price: 235,
        qty: 5,
        side: "SELL",
      },
      {
        id: "edidjt",
        price: 700,
        qty: 245,
        side: "BUY",
      },
      {
        id: "ed3djw",
        price: 236,
        qty: 245,
        side: "BUY",
      },
    ],
    ttl_phase_one: Date.now() + 10000,
    ttl_phase_two: 1000 * 60 * 60 * 24 * 31 * 6,
    funding_amount: 125_000,
    funding_date: null,

    timeseries: [
      { timestamp: 1668821699865, price: 567 },
      { timestamp: 1668821799865, price: 670 },
      { timestamp: 1668821899865, price: 589 },
      { timestamp: 1668821999865, price: 400 },
    ],
  };

  return { props: { security } };
}

interface Securities {
  security: Security;
}
interface Result {
  Id: string;
  Name: string;
}

export default function Security(props: Securities) {
  const [quantity, setQuantity] = useState(0);
  const [offer, setOffer] = useState(0);
  const { security } = props;
  const timeToNextPhase = security.funding_date
    ? formatDistanceToNow(security.funding_date + security.ttl_phase_two)
    : formatDistanceToNow(security.creation_date + security.ttl_phase_one);

  function handleOrder(action: string) {
    const order = {
      request: "add",
      security: security.security_id,
      qty: quantity,
      price: offer,
      side: action,
      user: "user_id_hardcoded",
    };
    // const result = await fetch("", { method: "POST"})
  }

  const pageNotSearching = (
    <section className="grid grid-rows-[0.5fr_2fr] grid-cols-2 gap-2">
      <section>
        <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl pb-4">
          {security.title}
        </h2>
        <p className="text-xl">{security.description}</p>
        <p className="text-xl">{security.creator.name}</p>
        <p className="text-xl">{timeToNextPhase} until the next phase!</p>
        <Setps />
      </section>
      <section className="row-span-2 p">
        <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl py-4">
          Orders
        </h2>
        {security.orders
          .sort((i, j) => j.price - i.price)
          .map((order) => {
            return (
              <section
                key={order.id}
                className={`max-w-lg  border shadow rounded my-2 p-4 flex justify-between ${
                  order.side === "BUY" ? "bg-green-300" : "bg-red-300"
                }`}
              >
                <p>
                  {order.qty} units @ {order.price / 1000}€
                </p>
              </section>
            );
          })}
      </section>
      <section>
        <Graph timeseries={security.timeseries} />
        <div className="flex">
          <div className="flex flex-col mr-5">
            Quantity
            <input
              className="
  form-control
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
              type="number"
              name="quantity"
              id="quantity"
              onChange={(e) => setQuantity(Number(e.target.value))}
              value={quantity}
            />
          </div>{" "}
          <div className="flex flex-col ml-5">
            Price{" "}
            <input
              className="
  form-control
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
              type="number" // number input ist immer integer, somit keine floats möglich => Textinput benutzen?
              name="offer"
              id="offer"
              onChange={(e) => setOffer(Number(e.target.value))}
              value={offer}
            />
          </div>{" "}
          {(quantity * offer) / 100 != 0
            ? "= " + (quantity * offer) / 100
            : " "}
        </div>
        <div className="flex gap-20 pt-5">
          <button
            onClick={() => handleOrder("BUY")}
            className="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
          >
            Buy Order
          </button>
          <button
            onClick={() => handleOrder("SELL")}
            type="button"
            className="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
          >
            Sell Order
          </button>
        </div>
      </section>
      <h2 className="text-4xl font-bold tracking-tight  sm:text-5xl py-4">
        Further information
      </h2>
    </section>
  );

  return (
    <>
      <Head>
        <title>{security.title}</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Layout>{pageNotSearching}</Layout>
    </>
  );
}
