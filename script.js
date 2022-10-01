const form = document.getElementById("form");
const con = document.getElementById("api");

const origin = document.getElementById("Origin");
const destination = document.getElementById("destination");


function initMap() {
    const google = window.google;
    var map = new google.maps.Map(document.getElementById('map'), {
      zoom: 7,
      center: {lat: 41.85, lng: -87.65}
    });
    var directionsService = new google.maps.DirectionsService;
    var directionsDisplay = new google.maps.DirectionsRenderer;
    directionsDisplay.setMap(map);

    var onSubmit = function(e) {
        var mapRequest = {
            origin: origin.value,
            destination: destination.value,
            travelMode: 'DRIVING'
        };
        
        e.preventDefault();
        calcAndDisplayRoute(directionsService, directionsDisplay, mapRequest);
      };

    var onSubmitapi = function(e) {
        console.log("inside")
        var mapRequest = {
            origin: origin.value,
            destination: destination.value,
            travelMode: 'DRIVING'
        };
        
        e.preventDefault();
        calcAndDisplayAPI(directionsDisplay,mapRequest);
      };


    var options = {
        types: ['(cities)']
    }
    var autocomplete1 = new google.maps.places.Autocomplete(origin, options);
    var autocomplete2 = new google.maps.places.Autocomplete(destination, options);
    con.addEventListener("submit", onSubmitapi); 
    form.addEventListener("submit", onSubmit)
    

}

//function to calculate 
async function calcAndDisplayRoute(directionsService, directionsDisplay, mapRequest) {
    directionsService.route(mapRequest, function (result, status) {
        if (status == google.maps.DirectionsStatus.OK) {

            //Get distance and time
            const output = document.querySelector('#output');
            output.innerHTML = `<div class="resultcontainer" >Driving distance: `+ result.routes[0].legs[0].distance.text + ".<br />Duration: " + result.routes[0].legs[0].duration.text + ".</div>";
            console.log(result)
            //display route
            directionsDisplay.setDirections(result);
        } else {
            //delete route from map
            
            //show error message
            output.innerHTML = ` <div class="noresultcontainer" style="color:white";> Could not retrieve driving distance.</div> `;
        }
    });
}

//function to get disrance and duration from api 
async function calcAndDisplayAPI(directionsDisplay,mapRequest){
    await fetch("https://goserver.gpejavf3ccgpf0c6.eastasia.azurecontainer.io:8080/getdirection", {
        method: 'POST',
        header: new Headers({
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "https://goserver.gpejavf3ccgpf0c6.eastasia.azurecontainer.io:8080/",
          }),
          mode: "cors",
        body: JSON.stringify({
            'origin': mapRequest.origin,
            'destination': mapRequest.destination
        })
    })
    .then(response => {
        return response.json();
    })
    .then(data => {
        directionsDisplay.setDirections({ routes: [] });
        console.log(data)
        output.innerHTML = `<div class="resultcontainer" >Driving distance: `+ data.distance + ".<br />Duration: " + data.duration + ".</div>";

    })
    .catch(errResData => {
        directionsDisplay.setDirections({ routes: [] });
        console.log(errResData)
        output.innerHTML = ` <div class="noresultcontainer" style="color:white";> Could not retrieve driving distance.</div> `;
    });
}

window.initMap = initMap;

