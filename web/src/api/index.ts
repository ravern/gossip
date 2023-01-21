import Axios, { AxiosRequestConfig } from "axios";
import { QueryClient } from "react-query";

import Config from "src/config";

const LOCAL_STORAGE_KEY_ACCESS_TOKEN = "accessToken";

export const axiosClient = Axios.create({
  baseURL: Config.api.baseURL,
});

axiosClient.defaults.headers.common["Content-Type"] = "application/json";
axiosClient.interceptors.request.use(
  (config: AxiosRequestConfig): AxiosRequestConfig => {
    config.headers["Authorization"] = `Bearer ${localStorage.getItem(
      LOCAL_STORAGE_KEY_ACCESS_TOKEN
    )}`;
    return config;
  }
);

export const queryClient = new QueryClient();
