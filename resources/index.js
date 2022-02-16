// create function to get data api weather pemalang city
const getData = async () => {
    const response = await fetch(
        "https://api.openweathermap.org/data/2.5/weather?q=Pemalang&appid=6a1c8a8e6e0b6a7b4c4b8b1c6a8f9a9b"
    );
    const data = await response.json();
    return data;
    };