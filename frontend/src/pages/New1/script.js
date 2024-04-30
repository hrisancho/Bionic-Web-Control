const uuidInput = document.getElementById('uuid');
const refreshButton = document.getElementById('refresh-button');
const strainGaugePanel = document.getElementById('strain-gauge-panel');
const servoPanel = document.getElementById('servo-panel');
const potentiometerPanel = document.getElementById('potentiometer-panel');

refreshButton.addEventListener('click', refreshData);

function refreshData() {
    const uuid = uuidInput.value;
    if (!uuid) {
        return;
    }

    fetchData('http://127.0.0.1:8080/api/v1/hand/123/monitoring/strain-gauge/all-finger', strainGaugePanel);
    // fetchData('http://127.0.0.1:8080/api/v1/hand/' + uuid + '/monitoring/servo/info/all-servo', servoPanel);
    // fetchData('http://127.0.0.1:8080/api/v1/hand/' + uuid + '/monitoring/potentiometer/all-potentiometer', potentiometerPanel);
}

function fetchData(url, panel) {
    fetch(url)
        .then(response => response.json())
        .then(data => {
            updatePanelContent(panel, data);
        })
        .catch(error => {
            console.error('Error fetching data:', error);
        });
}

function updatePanelContent(panel, data) {
    const contentElement = panel.querySelector('.data-content');

