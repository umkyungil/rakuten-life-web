function initMap() {
    var latlng = new google.maps.LatLng(43.066655,141.347556);
    var myOptions = {
        zoom: 17, /*拡大比率*/
        center: latlng, /*表示枠内の中心点*/
        mapTypeId: google.maps.MapTypeId.ROADMAP/*表示タイプの指定*/
    };
    var map = new google.maps.Map(document.getElementById('map_canvas'), myOptions);

    var markerOptions = {
        position: latlng,
        map: map,
    };
    var marker = new google.maps.Marker(markerOptions);

    google.maps.event.addDomListener(window, 'load', initMap);
    google.maps.event.addDomListener(window, 'resize', function(){
        map.panTo(latlng);//ウィンドウがリサイズされたら中心点に追従
    });

    $(".popup-modal").click(function(){
        google.maps.event.trigger(map, 'resize');
        map.setCenter(latlng);
    })//モーダルが表示されたら地図を再描画。display:none対応。
}