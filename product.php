<?php

class Product {
    //array of ['G01'=>['product'=>'Green','price'=>24.95]]
    private $products = [];
    //e.g. ['50'=>4.95,'90'=>2.95,'inf'=>0]
    private $shipping_fees = [];
    //e.g buy one red widget,get the second half price will be modeled as ['R01'=>[2=>'0.5']]
    private $offers = [];
    //e.g array of ['G01'=>2,'R01'=>3]
    private $basket = [];

    public function __construct($products,$shipping_fees,$offers=[]) {
        $this->products = $products;
        $this->shipping_fees = $shipping_fees;
        $this->offers = $offers;
    }

    public function addProducts($new_codes){
        $unique = array_unique($new_codes);
        foreach($unique as $code){
            $count = count(array_keys($new_codes,$code));
            $this->basket[$code] = $count;
        }
        return $this->basket;
    }

    public function getTotalPrice() {
        $sum = 0;
        $codes = $this->basket;
        //Iterating over all the products in a basket
        foreach($this->basket as $code=>$num) {
            //Check if the product exists in product list. 
            //If it doesn't exist, just skip.
            if(isset($this->products[$code])){
                //Price of the product whose code is $code
                $price = $this->products[$code]['price'];
                if(!isset($this->offers[$code]))//doesn't have an offer rule
                    $sum += $price * $num;
                else{// Have an offer rule.
                    for($i=0;$i<$num;$i++){
                        $found_offer = 0;//Flag for found offer rule for a given product
                        foreach($this->offers[$code] as $key=>$value){
                            if($i+1 == intval($key)){
                                $found_offer = 1;
                                //Apply percentage to the price
                                $sum += $price * $value;
                                break;
                            }
                        }
                        if($found_offer == 0)
                            $sum += $price; 
                    }
                }
            }
        }

        //Applying Shipping Charge Rules
        foreach($this->shipping_fees as $limit=>$fee){
            if($limit == 'inf'){
                $sum += $fee;
                break;
            }
            else if($sum < intval($limit)){
                $sum += $fee;
                break;
            } 
        }
        return $sum;
    }
}

$products = [
    'R01'=>[
        'product'=>'Red Widget',
        'price'=>32.95
    ],
    'G01'=>[
        'product'=>'Green Widget',
        'price'=>24.95
    ],
    'B01'=>[
        'product'=>'Blue Widget',
        'price'=>7.95
    ],
];

$basket = ['G01','R01','B01','G01','R01','R01'];
$shipping_fees = ['50'=>4.95,'90'=>2.95,'inf'=>0];
$offers = [ // This offers needs to be modeled in detail according to the categories.
    'G01'=>[
        '2'=>0.5,
        '3'=>0.4
    ],
    'R01'=>[
        '2'=>0.7
    ]
];
$product_object = new Product($products,$shipping_fees,$offers);
$product_object->addProducts($basket);

echo $product_object->getTotalPrice();