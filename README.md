# Laravel Socialite

This base will help to oauth2 login with Google, Facebook, Github

## Install

```shell
composer require viethqb/laravel-socialite
```

## Publish configuration file and Base Classes

```shell
php artisan vendor:publish --provider="Viethqb\LaravelSocialite\Providers\SocialiteServiceProvider"
```

## create file .env

```shell
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_REDIRECT=http://your-app-url/auth/google/callback

FACEBOOK_CLIENT_ID=your-facebook-client-id
FACEBOOK_CLIENT_SECRET=your-facebook-client-secret
FACEBOOK_REDIRECT=http://your-app-url/auth/facebook/callback

GITHUB_ID=your-github--client-id
GITHUB_SECRET=your-github-client-secret
GITHUB_REDIRECT=http://your-app-url/auth/github/callback
```

# Use to function base 

```shell
extends Base/SocialiteService.php

  public function redirectToProvider(SocialiteEnum $provider);

  public function handleProviderCallback(SocialiteEnum $provider);
```