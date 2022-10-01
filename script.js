const form = document.getElementById("form");
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

    var options = {
        types: ['(cities)']
    }
    var autocomplete1 = new google.maps.places.Autocomplete(origin, options);
    var autocomplete2 = new google.maps.places.Autocomplete(destination, options);
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

window.initMap = initMap;

