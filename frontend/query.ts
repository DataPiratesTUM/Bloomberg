import axios from "axios";
import { useQuery, useQueryClient, QueryClient } from "@tanstack/react-query";

export default function query() {
  // GET all history
  const getHistoryAll = () => {
    const config = {
      method: "get",
      url: "https://transaction.ban.app/order/history/security/all",
      headers: {},
    };
    return useQuery(["matchesAll"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  };

  // GET user history
  const getUserHistory = ({ userId }: { userId: string }) => {
    const config = {
      method: "get",
      url: "https://transaction.ban.app/order/value",
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
      url: "https://transaction.ban.app/order/history/security/",
      headers: {},
    };
    return useQuery(["portfolioValue"], async (): Promise<Array<any>> => {
      const { data } = await axios(config);
      return data;
    });
  }

  // POST buy order
  function postBuyOrder({ order, userId }: { order: Order; userId: string }) {
    const data = { order };
    const config = {
      method: "post",
      url: "https://transaction.ban.app/order/place",
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
      url: "https://transaction.ban.app/order/sell",
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
