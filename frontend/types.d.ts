declare interface Security {
  id: string;
  name: string;
  qty: number;
  price: number;
  price_bought: number;
}

declare type Timeseries = {
  timestamp: number;
  price: number;
}[];

declare interface User {
  id: string;
  name: string;
  balance: number;
  securities: Security[];
  timeseries: Timeseries;
}

declare interface Home {
  user: User;
}
