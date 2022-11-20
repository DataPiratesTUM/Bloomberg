declare type Timeseries = {
  timestamp: number;
  price: number;
}[];

declare interface TimeseriesCole {
  created: number;
  price: number;
  quantity: number;
}
declare interface OpenOrder {
  price: number;
  quantity: number;
  security: string;
  side: string;
}

declare interface Order {
  security: string;
  quantity: number;
  price: number;
  side: "buy" | "sell";
}

declare interface Portfolio {
  time: number;
  value: number;
}

declare interface Match {
  created: number;
  price: number;
  quantity: number;
  security: string;
}

declare interface Security {
  security_id: string;
  price: number;
  title: string;
  creator: string;
  description: string;
  fundingAmount: number;
  creationDate: number;
  fundingDate: number | null;
  ttl_phase_one: number;
  ttl_phase_two: number;
  price: number;
  quantity: number;
}

declare interface User {
  Name: string;
  OrganisationId: string;
  // user_id: string;
  // name: string;
  // balance: number;
  // securities: {
  //   id: string;
  //   name: string;
  //   qty: number;
  //   price: number;
  //   price_bought: number;
  // }[];
  // timeseries: Timeseries;
}

declare type Trending = string[];
declare interface SecurityOverview {
  Id: string;
  Name: string;
}
