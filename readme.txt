The Product Class has all the data and functions in it.
1. Variables
$products is a variable for product categories, e.g.
$products = [
    'R01'=>[
        'product'=>'Red Widget',
        'price'=>32.95
    ]
];
R01 is its code and the key for the detail of a product.

$shipping_fees is an array for delivery charge rules.
The keys of this array is upper limit of the shipping condition.
For example, "Orders under $50 cost $4.95" can be modeled as ["50"=>4.95]
The assumption is that the keys of this array will increase sequentially, not just a random array.
$offers is an array for offer rules.

"buy one red widget,get the second half price" can be modeled as 
"R01"=>[ 
    '2'=>0.5
]
R01 is the code for a product and 2 stands for the sequence of a product that will get a benefit.
0.5 stands for the price percentage of the benefited product.
Of course, this model can't fully cover all the offer rules.
This is just for demonstration purpose.
If many offer rules are given , then I will need to generalize them and 
extract standardized model from it.

$baskets is an array that consists of product codes in basket.

2. Methods
For running this app, first the object of Product class will be initialized with $products, 
$shipping_fees, and $offeres.
-Constructor
These values will be assigned to member variables of an object using constructor.
-addProducts
This function will count each product among the products in a basket.
As a result, the $basket member variable will be assigned by this value.
-getTotalPrice
This is the function that calculates the total price of products in a basket 
by taking into account the shipping charge rules and offers.
The details are commented in the code.

Thanks for your time to read my code.
