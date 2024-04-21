import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from '../components/ui/Navbar';
import SettingsForm from '../components/ui/forms/SettingsForm';

const SettingsPage: React.FC = () => {
    return (
        <div className='bg-green-500 h-screen w-screen '>
            <Navbar />
            <div className="flex mt-16 justify-center items-center">
                <SettingsForm />
            </div>
        </div>
    );
}

export default SettingsPage;
