import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";

import SuccessBtn from "../buttons/SuccessBtn";
import FormField from "../fields/FormField";

import { AppDispatch } from "../../../services/state/store";
import { signUpAsync } from "../../../services/state/user/userSlice";

const SignUpForm: React.FC = () => {
  const [name, setName] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const dispatch = useDispatch<AppDispatch>();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    console.log("Name:", name);
    console.log("Username:", username);
    console.log("Email:", email);
    console.log("Password:", password);

    dispatch(
      signUpAsync({
        name,
        username,
        email,
        password,
      })
    );
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
        <label className="text-gray-700 font-bold mb-2">Sign Up</label>
      </div>
      <div className="mb-4">
        <FormField
          id="name"
          placeholder="Name"
          value={name}
          onChange={(e: any) => setName(e.target.value)}
        />
      </div>
      <div className="mb-4">
        <FormField
          id="username"
          placeholder="Username"
          value={username}
          onChange={(e: any) => setUsername(e.target.value)}
        />
      </div>
      <div className="mb-4">
        <FormField
          id="email"
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e: any) => setEmail(e.target.value)}
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
          Sign Up
        </SuccessBtn>
      </div>
    </form>
  );
};

export default SignUpForm;
