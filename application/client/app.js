'use strict';

var app = angular.module('application', []);

app.controller('AppCtrl', function($scope, appFactory) {
    // 사용자 등록 함수
    $scope.registerUser = function() {
        var user = {
            userID: $scope.user.userID,
            name: $scope.user.name,
            buildingName: $scope.user.buildingName,
            description: $scope.user.description
        };
        appFactory.registerUser(user, function(data) {
            $scope.registerUserResult = data;
        });
    };

    // 부동산 조회 함수
    $scope.viewBuilding = function() {
        appFactory.viewBuilding($scope.buildingName, function(data) {
            $scope.buildingInfo = data;
        });
    };

    // JSON 데이터를 한 줄씩 내려서 표시하는 함수
    $scope.formatBuildingInfo = function(buildingInfo) {
        if (!buildingInfo) return ''; // 데이터가 없을 경우 빈 문자열 반환

        // JSON 데이터를 파싱하여 객체로 변환
        var buildingObj = JSON.parse(buildingInfo);

        // 객체의 각 속성을 반복하면서 속성 이름과 값을 새 줄에 하나씩 표시
        var formattedInfo = '';
        for (var key in buildingObj) {
            formattedInfo += key + ': ' + buildingObj[key] + '\n';
        }

        return formattedInfo; // 포맷팅된 문자열 반환
    };
});


app.factory('appFactory', function($http) {
    var factory = {};

    // 사용자 등록을 위한 HTTP POST 요청 함수
    factory.registerUser = function(user, callback) {
        $http.post('/registerUser', user).then(function(response) {
            callback(response.data);
        }, function(error) {
            callback({ error: error });
        });
    };

    // 부동산 조회를 위한 HTTP GET 요청 함수
    factory.viewBuilding = function(buildingName, callback) {
        $http.get('/viewBuilding?buildingName=' + buildingName).then(function(response) {
            callback(response.data);
        }, function(error) {
            callback({ error: error });
        });
    };

    return factory;
});
