import React  from 'react';
import ReactWeather, { useOpenWeather } from 'react-open-weather';
import './App.css';

function App() {
  const { data, isLoading, errorMessage } = useOpenWeather({
    key: process.env.REACT_APP_WEATHER_API_KEY,
    lat: '1.2921',
    lon: '36.8219',
    lang: 'en',
    unit: 'metric', // values are (metric, standard, imperial)
  });

  return (
    <div className="App">
      <ReactWeather
        isLoading={isLoading}
        errorMessage={errorMessage}
        data={data}
        lang="en"
        locationLabel="Nairobi"
        unitsLabels={{ temperature: 'C', windSpeed: 'Km/h' }}
        showForecast
      />
      <h1> Hello world </h1>
      <p> I'm a paragraph </p>
    </div>
  );
}

export default App;
