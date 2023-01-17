import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import { Text, ChakraProvider } from '@chakra-ui/react'
import icon from '../../assets/icon.svg';
import './App.css';

const Hello = () => {
  return(<h1>Abc</h1>)
}

export default function App() {
  return (
    <ChakraProvider>
    <Router>
      <Routes>
        <Route path="/" element={<Hello />} />
      </Routes>
    </Router>
    </ChakraProvider>
  );
}
