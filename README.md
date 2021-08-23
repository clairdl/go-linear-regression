# Regression in go

### Intro

Regression is the process of statistically modelling the relationship between _n_ variables from known data points, which enables us to predict unknown values. Regression techniques are generally concerned with predicting continous values, as opposed to a discreet set of categories.

### Linear regression

Linear regression is the simplest but possibly the most fundamental model, and uses the good old: 

`y = mx + b`

Which describes a line with slope `m` and y-intercept `b`.

One way of calculating `m` and `b` is called the 'ordinary least squares' method:

1. Randomise values for both m and b to create an example line
2. Find the distance between the example line and each value in the dataset. These distances are called 'errors':

<img src="./docs/ols.png">

3. Sum the squares of these errors:

<p align="center">
<img src="./docs/eq0.png">
</p>

Now, we iteratively adjust the values of `m` and `b` in order to minimize this sum. A popular and general optimization technique to find local minima is called _gradient descent_, but that's a topic for another day :)

The accuracy and performance of linear regression is dependant on its **assumptions**:

- **Linearity**: there is a linear relationship between the dependant variable and the independant variable(s)
- **Normality**: your variables are distributed normally
- **No multicollinearity**: your independant variables should not be predictors of eachother
- **No auto-correlation**: a fancy way of saying your variables should not depend on themselves, i.e they are not values in a time series, e.g. the S&P 500, for example

**Pitfalls**:
- Extrapolation beyond the model can quickly become very inaccurate
- Extreme outliers can throw off the model 


