import { PayloadAction, createAsyncThunk, createSlice } from "@reduxjs/toolkit";

import User from "../../../models/User";
import { API_HOST } from "../../../constants";

const API_LOGIN_URL = API_HOST + "/api/v1/auth/signin";
const API_REGISTER_URL = API_HOST + "/api/v1/auth/signup";

interface UserState {
  user: User | null;
  status: "idle" | "loading" | "succeeded" | "failed";
  error: string | null;
}

const initialState: UserState = {
  user: null,
  status: "idle",
  error: null,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getByTokenAsync.pending, (state) => {
        state.status = "loading";
        console.log("Pending user");
      })
      .addCase(
        getByTokenAsync.fulfilled,
        (state, action: PayloadAction<User | null | undefined>) => {
          state.user = action.payload!;
          state.status = "succeeded";
          console.log("User loaded");
        }
      )
      .addCase(getByTokenAsync.rejected, (state, action) => {
        state.error = action.error.message!;
        state.status = "failed";
        console.log("User failed");
      });
    builder
      .addCase(signUpAsync.pending, (state) => {
        state.status = "loading";
        console.log("Pending user");
      })
      .addCase(
        signUpAsync.fulfilled,
        (state, action: PayloadAction<User | null | undefined>) => {
          state.user = action.payload!;
          state.status = "succeeded";
          console.log("User loaded");
        }
      )
      .addCase(signUpAsync.rejected, (state, action) => {
        state.error = action.error.message!;
        state.status = "failed";
        console.log("User failed");
      });
    builder
      .addCase(signInAsync.pending, (state) => {
        state.status = "loading";
        console.log("Pending user");
      })
      .addCase(
        signInAsync.fulfilled,
        (state, action: PayloadAction<User | null | undefined>) => {
          state.user = action.payload!;
          state.status = "succeeded";
          console.log("User loaded");
        }
      )
      .addCase(signInAsync.rejected, (state, action) => {
        state.error = action.error.message!;
        state.status = "failed";
        console.log("User failed");
      });
    builder
      .addCase(logoutAsync.pending, (state) => {
        state.status = "loading";
        console.log("Logout Pending");
      })
      .addCase(
        logoutAsync.fulfilled,
        (state, action: PayloadAction<User | null | undefined>) => {
          state.user = action.payload!;
          state.status = "succeeded";
          console.log("Logout done");
        }
      )
      .addCase(logoutAsync.rejected, (state, action) => {
        state.error = action.error.message!;
        state.status = "failed";
        console.log("Logout failed");
      });
  },
});

export const getByTokenAsync = createAsyncThunk(
  "user/getByTokenAsync",
  async () => {
    console.log("Start loading user...");
    const requestOptions: RequestInit = {
      method: "GET",
      credentials: "include",
    };

    try {
      const response = await fetch(
        `${API_HOST}/api/v1/auth/me`,
        requestOptions
      );
      if (response.ok) {
        const data: User = await response.json();
        return data;
      } else {
        console.error("Failed to fetch user");
        return null;
      }
    } catch (err) {
      console.error("Error fetching user:", err);
      return null;
    }
  }
);

export const signUpAsync = createAsyncThunk(
  "user/signUpAsync",
  async ({
    name,
    username,
    email,
    password,
  }: {
    name: string;
    username: string;
    email: string;
    password: string;
  }) => {
    console.log("Start Sign Up...");
    const requestOptions: RequestInit = {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name,
        username: username,
        email: email,
        password: password,
      }),
    };

    try {
      const response = await fetch(API_REGISTER_URL, requestOptions);
      if (response.ok) {
        const data: User = await response.json();
        return data;
      } else {
        console.error("Failed to fetch user");
        return null;
      }
    } catch (err) {
      console.error("Error fetching user:", err);
      return null;
    }
  }
);

export const signInAsync = createAsyncThunk(
  "user/signInAsync",
  async ({ username, password }: { username: string; password: string }) => {
    console.log("Start Sign In...");
    const requestOptions: RequestInit = {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    };

    try {
      const response = await fetch(API_LOGIN_URL, requestOptions);
      if (response.ok) {
        const data: User = await response.json();
        return data;
      } else {
        console.error("Failed to fetch user");
        return null;
      }
    } catch (err) {
      console.error("Error fetching user:", err);
      return null;
    }
  }
);

export const logoutAsync = createAsyncThunk("user/logoutAsync", async () => {
  console.log("Start Log Out...");
  const requestOptions: RequestInit = {
    method: "DELETE",
    credentials: "include",
  };

  try {
    const response = await fetch(
      `${API_HOST}/api/v1/auth/logout`,
      requestOptions
    );
    if (response.ok) {
      return null;
    } else {
      console.error("Failed to fetch user");
      return null;
    }
  } catch (err) {
    console.error("Error fetching user:", err);
    return null;
  }
});

export default userSlice.reducer;
