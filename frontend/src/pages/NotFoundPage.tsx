import React from "react";
import { Link } from "react-router-dom";

import SuccessBtn from "../components/ui/buttons/SuccessBtn";

const NotFoundPage: React.FC = () => {
  return (
    <div className="flex items-center justify-center h-screen bg-gray-200">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-gray-800">404</h1>
        <h2 className="text-2xl font-semibold text-gray-700 mb-4">
          Страница не найдена
        </h2>
        <p className="text-lg text-gray-600">
          Возможно, страница, которую вы ищете, удалена или временно недоступна.
        </p>
        <Link to="/">
          <SuccessBtn className="mt-6 h-10">Вернуться на Главную</SuccessBtn>
        </Link>
      </div>
    </div>
  );
};

export default NotFoundPage;
