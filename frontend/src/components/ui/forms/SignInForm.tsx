import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";

import FormField from "../fields/FormField";
import SuccessBtn from "../buttons/SuccessBtn";

import { AppDispatch } from "../../../services/state/store";
import { signInAsync } from "../../../services/state/user/userSlice";

const SignInForm: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const dispatch = useDispatch<AppDispatch>();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    dispatch(signInAsync({ username, password }));
    navigate("/");
  };

  return (
    <form
      className="bg-white shadow-md rounded px-32 pt-16 pb-16"
      onSubmit={handleSubmit}
    >
      <div className="mb-4 flex justify-center">
        <img src="/assets/Logo.svg" alt="Logo" className="rounded-lg" />
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
