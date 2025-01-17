import axios from "axios";

const API = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/v1",
  withCredentials: true,
});

let isRefreshing = false;
let refreshSubscribers: ((token: string) => void)[] = [];

const TokenService = {
  getToken: () => localStorage.getItem("token"),
  setToken: (token: string) => localStorage.setItem("token", token),
  removeToken: () => localStorage.removeItem("token"),
};

const onRefreshed = (token: string) => {
  refreshSubscribers.forEach((callback) => callback(token));
  refreshSubscribers = [];
};

const addRefreshSubscriber = (callback: (token: string) => void) => {
  refreshSubscribers.push(callback);
};

API.interceptors.request.use(
  (config) => {
    const token = TokenService.getToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error),
);

API.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (error.response) {
      if (error.response.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true;
        if (isRefreshing) {
          return new Promise((resolve) => {
            addRefreshSubscriber((token) => {
              originalRequest.headers.Authorization = `Bearer ${token}`;
              resolve(API(originalRequest));
            });
          });
        }

        isRefreshing = true;

        try {
          const { data } = await axios.post(
            `${API.defaults.baseURL}/refresh`,
            {},
            { withCredentials: true },
          );

          const { accessToken } = data;
          TokenService.setToken(accessToken);
          API.defaults.headers.common["Authorization"] =
            `Bearer ${accessToken}`;
          onRefreshed(accessToken);
          isRefreshing = false;
          return API(originalRequest);
        } catch {
          TokenService.removeToken();
          isRefreshing = false;
          window.location.href = "/login";
        }
      }

      return Promise.reject(error);
    } else if (error.request) {
      return Promise.reject({
        message: "No response from server. Please check your connection.",
      });
    } else {
      return Promise.reject(error);
    }
  },
);

export default API;

