# Regression in Go

breifly, regression is the process of analysing and statistically modelling the relationship between variables (by drawing a curve through the data points), and predicting new values.

nb: regression techniques are concerned with predicting continous values, as opposed to a discreet set of categories.

_linear regression_, then, is a simple yet fundamental model:

we model the dependant variable (DV) y by an independant variable (IV) x using the following equation:

## y = mx + b

where m is the slope of the line and b is the y intercept

we know x and y from our dataset, so how do we find m and b? the most common method is called 'ordinary least squares' (OLS):

1. choose, or randomise values for both m and b, to create an example line
2. measure the vertical distance between each data point and the example line. these distances are called 'errors'

<img src="./docs/ols.png">

3. next, we sum the squares of these errors:

\frac{error_1^2+error_2^2+\dotsb+error_n^2}{n} 

now, we adjust (train) the values of m and b until we minimize this sum!

this process, of iteratively optimizing some value or parameter is called _gradient descent_, and is a ubiquitous technique within machine learning, including _deep learning_. 

let's think about gradient descent geometrically - imagine a point on a turbulent geometric shape in 3d space:

<img src="./docs/gradient-descent.jpg">

To reach the local minimum value algorithmically, you must steps toward a downward direction. but what if the topology has descents of varying steepness? how do you know in what direction to step toward? this problems 


### Linear Regression assumptions
like all machine learning models, linear regression doesn't work in all situations, and its performance is based its assumptions. you can basically think of these assumptions as 'requirements' for linear regression to work effectively:

- **Linearity:** there is a linear relationship between the dependant variable and the independant variable
- **Normality:** your variables are distributed normally - think bell curves
- **No multicollinearity:** your independant variables should not be able to be predicted by eachother
- **No auto-correlation:** Auto-correlation means that a variable depends on itself or some 
- **Homoscedasticity:** is an unnecessarily flashy word that means the errors (refer to figure 1) have a constant variance for each value of the independant variable


other things to keep in mind:
- Extrapolation is super unpredictable
- Avoid spurious correlations by reasoning for why two variables may be functionally related
- Outliers may throw off regression lines

















