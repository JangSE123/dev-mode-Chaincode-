'use strict';

var app = angular.module('application', []);

app.controller('AppCtrl', function($scope, appFactory){
   $("#success_init").hide();
   $("#success_qurey").hide();
   $scope.initAB = function(){
       appFactory.initAB($scope.abstore, function(data){
           if(data == "")
           $scope.init_ab = "success";
           $("#success_init").show();
       });
   }
   $scope.queryAB = function(){
       appFactory.queryAB($scope.walletid, function(data){
           $scope.query_ab = data;
           $("#success_qurey").show();
       });
   }
});
app.factory('appFactory', function($http){
      
    var factory = {};
 
    factory.initAB = function(data, callback){
        $http.get('/init?a='+data.a+'&aval='+data.aval+'&b='+data.b+'&bval='+data.bval+'&c='+data.c+'&cval='+data.cval).success(function(output){
            callback(output)
        });
    }
    factory.queryAB = function(name, callback){
        $http.get('/query?name='+name).success(function(output){
            callback(output)
        });
    }
    return factory;
 });
 // AppCtrl 컨트롤러 수정
app.controller('AppCtrl', function($scope, appFactory){
    // 기존 함수 유지
    $scope.initAB = function(){
        appFactory.initAB($scope.abstore, function(data){
            $scope.init_ab = "Initialization Success";
            $("#success_init").show();
        });
    };
    $scope.queryAB = function(){
        appFactory.queryAB($scope.walletid, function(data){
            $scope.query_ab = data;
            $("#success_qurey").show();
        });
    };

    // 사용자 등록 함수 추가
    $scope.registerUser = function(){
        appFactory.registerUser($scope.user, function(data){
            $scope.register_user_result = data;
        });
    };

    // 부동산 등록 함수 추가
    $scope.registerProperty = function(){
        appFactory.registerProperty($scope.property, function(data){
            $scope.register_property_result = data;
        });
    };
});

// appFactory 팩토리 수정
app.factory('appFactory', function($http){
    var factory = {};

    // 기존 함수 유지
    factory.initAB = function(data, callback){
        $http.get('/init?a='+data.a+'&aval='+data.aval+'&b='+data.b+'&bval='+data.bval+'&c='+data.c+'&cval='+data.cval).success(function(output){
            callback(output);
        });
    };
    factory.queryAB = function(name, callback){
        $http.get('/query?name='+name).success(function(output){
            callback(output);
        });
    };

    // 사용자 등록을 위한 HTTP POST 요청 함수 추가
    factory.registerUser = function(user, callback){
        $http.post('/registerUser', user).success(function(output){
            callback(output);
        });
    };

    // 부동산 등록을 위한 HTTP POST 요청 함수 추가
    factory.registerProperty = function(property, callback){
        $http.post('/registerProperty', property).success(function(output){
            callback(output);
        });
    };

    return factory;
});

 