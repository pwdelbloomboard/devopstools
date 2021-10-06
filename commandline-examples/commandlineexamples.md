# Using Grep

If you want to know whether a particular feature is already included in the codebase, you can use grep to search for things.

For example, if you want to know whether brew is included within an installation script, you can use:

```
grep -ri brew *
```

This will perform a greedy search across the codebase for the term: brew

The “r” flag is recursive, across multiple files, while the “i" flag is inclusive

Using quotes brings up different results:

```
grep -ri ‘brew’ * 
```

Basically it’s searching directly for that term.  So as an example search, if you ran the above command within a codebase that includes a file called, "setup" to see if said file includes what it may pull up during the current iteration is something on the order of the following:

```
setup:        if ! command -v brew &>/dev/null; then
setup:            echo "Installing homebrew."
setup:            /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
setup:        BREW_TAPS=(homebrew/cask homebrew/services mongodb/brew)
setup:        CURRENT_TAPS=$(brew tap)
setup:                brew tap "$TAP"
setup:        for DIRECTORY in homebrew-core homebrew-cask
setup:            FULL_PATH=/usr/local/Homebrew/Library/Taps/homebrew/$DIRECTORY
setup:        CURRENT_PACKAGES=$(brew list --formula)
setup:                brew install "$PACKAGE"
setup:            CURRENT_CASKS=$(brew list --casks)
setup:                    brew install --cask "$PACKAGE"
setup:        DNSMASQ_CONF="$(brew --prefix)/etc/dnsmasq.conf"
setup:        DNSMASQ_RESTART="sudo brew services restart dnsmasq"
```

Basically the above shows how brew is being installed.

# Environmental Variables

On a MacOS environmental variables can be displayed with:

```
env
```

MacOS used to use the bash shell so the way to display hard-coded environment variables was:

```
cat ~/.bash_profile
```
However MacOS upgraded to zsh in more recent OS versions so the way to see the zsh env is (within iTerm2):

```
cat ~/.zshrc
```
Of course, editing this could cause a lot of problems depending upon how the machine is set up and configured and there may be a customization warning.

With this warning, and during setup, there may be an alternate location to put zsh customizations.

```
$HOME/.custom-ohmyzsh     if you want to change any settings for Oh My ZSH
# $HOME/.custom-zsh         any custom aliases/functions/settings/whatever go here
```

## Setting an env

```
export VARIABLE=something
```
The above will make it so that when you enter in "echo $VARIABLE" you will get:

```
something
```

To add this permanently into the zsh, you use:

```
echo 'export VARIABLE=something' >> ~/.zshenv
```