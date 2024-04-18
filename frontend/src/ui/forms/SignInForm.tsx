import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import FormField from "../fields/FormField";
import SuccessBtn from "../buttons/SuccessBtn";

const API_HOST = "http://127.0.0.1:8000";
const API_LOGIN_URL = API_HOST + "/api/v1/auth/signin";

const SignInForm: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
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

    fetch(API_LOGIN_URL, requestOptions)
      .then((response) => {
        console.log("Status:", response.status);
        if (response.status === 200) {
          console.log("Authentication successful");
          // const from = location.state?.from || '/';
          // navigate(from);
          navigate("/workspaces");
        } else {
          console.log("Authentication failed");
        }
        return response.json();
      })
      .catch((error) => {
        console.log("error", error);
      });
  };

  return (
    <form
      className="bg-white shadow-md rounded px-32 pt-16 pb-16"
      onSubmit={handleSubmit}
    >
      <div className="mb-4 flex justify-center">
        <img src="/assets/logo.png" alt="Logo" className="w-24 h-24 rounded-full"/>
      </div>
      <div className="mb-4 flex justify-center">
        <label className="text-gray-700 font-bold mb-2">Sign In</label>
      </div>
      <div className="mb-4">
        <FormField
          id="username"
          placeholder="Username"
          value={username}
          onChange={(e: any) => setUsername(e.target.value)}
        />
      </div>
      <div className="mb-6">
        <FormField
          id="password"
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e: any) => setPassword(e.target.value)}
        />
      </div>
      <div>
        <SuccessBtn className="w-full" type="submit">
          Sign In
        </SuccessBtn>
      </div>
    </form>
  );
};

export default SignInForm;
