import axios from "axios";
import { useQuery, useQueryClient, QueryClient } from "@tanstack/react-query";

export default function query() {
  // GET all history
  const getHistoryAll = () => {
    const config = {
      method: "get",
      url: "localhost:3001/order/history/security/all",
      headers: {},
    };
    return useQuery(["historyAll"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  };

  // GET user history
  const getUserHistory = ({ userId }: { userId: string }) => {
    const config = {
      method: "get",
      url: "localhost:3001/order/value",
      headers: {
        "X-User-Id": userId,
      },
    };
    return useQuery(["userHistory"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  };

  // GET profolio value
  function getPorfolioValue() {
    const config = {
      method: "get",
      url: "localhost:3001/order/history/security/",
      headers: {},
    };
    return useQuery(["portfolio"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  }

  // POST buy order
  function postBuyOrder({ order, userId }: { order: Order; userId: string }) {
    const data = { order };
    const config = {
      method: "post",
      url: "localhost:3001/order/place",
      headers: {
        "X-User-Id": userId,
      },
      data: data,
    };
    return useQuery(["buyOrder"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  }

  // POST sell order
  function postSellOrder({ order, userId }: { order: Order; userId: string }) {
    const data = { order };
    const config = {
      method: "post",
      url: "localhost:3001/order/sell",
      headers: {
        "X-User-Id": userId,
      },
      data: data,
    };
    return useQuery(["sellOrder"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  }
}